package app

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/console"
	kldns "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/dns"
	"math/rand"
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/miekg/dns"
	"go.uber.org/fx"
	"kloudlite.io/apps/dns/internal/app/graph"
	"kloudlite.io/apps/dns/internal/app/graph/generated"
	"kloudlite.io/apps/dns/internal/domain"
	"kloudlite.io/common"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/config"
	httpServer "kloudlite.io/pkg/http-server"
	"kloudlite.io/pkg/repos"
)

type Env struct {
	CookieDomain        string `env:"COOKIE_DOMAIN"`
	EdgeCnameBaseDomain string `env:"EDGE_CNAME_BASE_DOMAIN" required:"true"`
	DNSDomainNames      string `env:"DNS_DOMAIN_NAMES" required:"true"`
}

type DNSHandler struct {
	domain              domain.Domain
	EdgeCnameBaseDomain string
	dnsDomainNames      []string
}

func (h *DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	msg.Answer = []dns.RR{}
	for _, q := range r.Question {
		switch q.Qtype {
		case dns.TypeNS:
			for _, name := range h.dnsDomainNames {
				rr := &dns.NS{
					Hdr: dns.RR_Header{
						Name:   q.Name,
						Rrtype: dns.TypeNS,
						Class:  q.Qclass,
						Ttl:    60,
					},
					Ns: fmt.Sprintf("%s.", name),
				}
				msg.Answer = append(msg.Answer, rr)
			}
		case dns.TypeA:
			msg.Authoritative = true
			d := q.Name
			todo := context.TODO()
			host := strings.ToLower(d[:len(d)-1])
			if strings.HasSuffix(host, h.EdgeCnameBaseDomain) {
				ips, err := h.domain.GetNodeIps(todo, nil)
				if err != nil || len(ips) == 0 {
					msg.Answer = append(msg.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: d, Rrtype: dns.TypeA, Class: dns.ClassINET},
					})
				}
				for _, ip := range ips {
					msg.Answer = append(msg.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: d, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300},
						A:   net.ParseIP(ip),
					},
					)
				}
				break
			} else {
				record, err := h.domain.GetRecord(todo, host)
				if err != nil || record == nil {
					msg.Answer = append(msg.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: d, Rrtype: dns.TypeA, Class: dns.ClassINET},
					})
				}
				rand.Shuffle(len(record.Answers), func(i, j int) {
					record.Answers[i], record.Answers[j] = record.Answers[j], record.Answers[i]
				})
				for _, a := range record.Answers {
					msg.Answer = append(msg.Answer, &dns.A{
						Hdr: dns.RR_Header{Name: d, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(30)},
						A:   net.ParseIP(a),
					},
					)
				}
			}
		}
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("HERE3", msg.Answer)
}

type ConsoleClientConnection *grpc.ClientConn

var Module = fx.Module(
	"app",
	config.EnvFx[Env](),
	repos.NewFxMongoRepo[*domain.Record]("records", "rec", domain.RecordIndexes),
	repos.NewFxMongoRepo[*domain.Site]("sites", "site", domain.SiteIndexes),
	repos.NewFxMongoRepo[*domain.AccountCName]("account_cnames", "dns", domain.AccountCNameIndexes),
	repos.NewFxMongoRepo[*domain.NodeIps]("node_ips", "nips", domain.NodeIpIndexes),
	cache.NewFxRepo[[]*domain.Record](),
	domain.Module,

	fx.Provide(func(conn ConsoleClientConnection) console.ConsoleClient {
		return console.NewConsoleClient((*grpc.ClientConn)(conn))
	}),

	fx.Invoke(func(lifecycle fx.Lifecycle, s *dns.Server, d domain.Domain, recCache cache.Repo[[]*domain.Record], env *Env) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				s.Handler = &DNSHandler{
					domain:         d,
					dnsDomainNames: strings.Split(env.DNSDomainNames, ","),
				}
				return nil
			},
		})
	}),
	fx.Provide(fxDNSGrpcServer),
	fx.Invoke(func(server *grpc.Server, dnsServer kldns.DNSServer) {
		kldns.RegisterDNSServer(server, dnsServer)
	}),
	fx.Invoke(func(
		server *fiber.App,
		d domain.Domain,
		env *Env,
		cacheClient cache.Client,
	) {
		schema := generated.NewExecutableSchema(
			generated.Config{Resolvers: graph.NewResolver(d)},
		)

		server.Post("/upsert-domain", func(c *fiber.Ctx) error {
			var data struct {
				Domain   string
				ARecords []string
			}
			err := c.BodyParser(&data)
			if err != nil {
				return err
			}
			err = d.UpsertARecords(c.Context(), data.Domain, data.ARecords)
			if err != nil {
				return err
			}
			c.Send([]byte("done"))
			return nil

		})
		server.Post("/upsert-node-ips", func(c *fiber.Ctx) error {
			var regionIps struct {
				Region string   `json:"region"`
				Ips    []string `json:"ips"`
			}
			err := c.BodyParser(&regionIps)
			if err != nil {
				return err
			}
			done := d.UpdateNodeIPs(c.Context(), regionIps.Region, regionIps.Ips)
			if !done {
				return fmt.Errorf("failed to update node ips")
			}
			c.Send([]byte("done"))
			return nil
		})
		server.Get("/get-records/:domain_name", func(c *fiber.Ctx) error {
			domainName := c.Params("domain_name")
			record, err := d.GetRecord(c.Context(), domainName)
			if err != nil {
				return err
			}
			r, err := json.Marshal(record)
			if err != nil {
				return err
			}
			c.Send(r)
			return nil
		})
		server.Delete("/delete-domain/:domain_name", func(c *fiber.Ctx) error {
			domainName := c.Params("domain_name")

			err := d.DeleteRecords(c.Context(), domainName)

			if err != nil {
				return err
			}

			c.Send([]byte("done"))
			return nil
		})

		httpServer.SetupGQLServer(
			server,
			schema,
			httpServer.NewSessionMiddleware[*common.AuthSession](
				cacheClient,
				common.CookieName,
				env.CookieDomain,
				common.CacheSessionPrefix,
			),
		)
	}),
)
