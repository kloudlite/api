package domain

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
	"go.uber.org/fx"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/errors"
	"kloudlite.io/pkg/repos"
)

type domainI struct {
	recordsRepo       repos.DbRepo[*Record]
	sitesRepo         repos.DbRepo[*Site]
	recordsCache      cache.Repo[[]*Record]
	accountCNamesRepo repos.DbRepo[*AccountCName]
}

func (d *domainI) GetSiteFromDomain(ctx context.Context, domain string) (*Site, error) {
	one, err := d.sitesRepo.FindOne(ctx, repos.Filter{
		"host": domain,
	})
	if err != nil {
		return nil, err
	}
	if one == nil {
		return nil, errors.New("site not found")
	}
	return one, nil
}

func (d *domainI) GetSites(ctx context.Context, accountId string) ([]*Site, error) {
	return d.sitesRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			"accountId": accountId,
		},
	})
}

func (d *domainI) CreateSite(ctx context.Context, domain string, accountId repos.ID) error {
	one, err := d.sitesRepo.FindOne(ctx, repos.Filter{
		"domain":    domain,
		"accountId": accountId,
	})
	if err != nil {
		return err
	}
	if one != nil {
		return errors.New("site already exists")
	}
	if one == nil {
		one, err = d.sitesRepo.Create(ctx, &Site{
			Domain:    domain,
			AccountId: accountId,
			Verified:  false,
		})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (d *domainI) VerifySite(ctx context.Context, siteId repos.ID) error {
	site, err := d.sitesRepo.FindById(ctx, siteId)
	if err != nil {
		return err
	}
	if site == nil {
		return errors.New("site not found")
	}
	if site.Verified {
		return errors.New("site already verified")
	}
	cname, err := net.LookupCNAME(site.Domain)
	if err != nil {
		return err
	}
	accountCnameIdentity, err := d.getAccountCName(ctx, string(site.AccountId))
	if err != nil {
		return err
	}
	if cname != fmt.Sprintf("%s.edgenet.khost.dev", accountCnameIdentity) {
		return errors.New("cname does not match")
	}
	err = d.sitesRepo.UpdateMany(ctx, repos.Filter{
		"host": site.Domain,
	}, map[string]any{
		"verified": false,
	})
	if err != nil {
		return err
	}
	site.Verified = true
	_, err = d.sitesRepo.UpdateById(ctx, site.Id, site)
	return err
}

func (d *domainI) GetSite(ctx context.Context, siteId string) (*Site, error) {
	return d.sitesRepo.FindById(ctx, repos.ID(siteId))
}

func (d *domainI) GetAccountEdgeCName(ctx context.Context, accountId string) (string, error) {
	name, err := d.getAccountCName(ctx, accountId)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.edgenet.khost.dev", name), nil
}

func (d *domainI) getAccountCName(ctx context.Context, accountId string) (string, error) {
	accountDNS, err := d.accountCNamesRepo.FindOne(ctx, repos.Filter{
		"accountId": accountId,
	})
	if err != nil {
		return "", err
	}
	if accountDNS == nil {
		seed := time.Now().UTC().UnixNano()
		nameGenerator := namegenerator.NewNameGenerator(seed)
		name1 := nameGenerator.Generate()
		name2 := nameGenerator.Generate()
		create, err := d.accountCNamesRepo.Create(ctx, &AccountCName{
			AccountId: repos.ID(accountId),
			CName:     fmt.Sprintf("%s-%s", name1, name2),
		})
		if err != nil {
			return "", err
		}
		return create.CName, nil
	}
	return accountDNS.CName, nil
}

func (d *domainI) GetRecords(ctx context.Context, host string) ([]*Record, error) {

	if matchedRecords, err := d.recordsCache.Get(ctx, host); err == nil && matchedRecords != nil {
		return matchedRecords, nil
	}

	domainSplits := strings.Split(strings.TrimSpace(host), ".")
	filters := make([]repos.Filter, 0)
	for i := range domainSplits {
		filters = append(filters, repos.Filter{
			"host": strings.Join(domainSplits[i:], "."),
		})
	}
	matchedSites, err := d.sitesRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			"verified": true,
			"$or":      filters,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(matchedSites) == 0 {
		return nil, errors.New("NoSitesFound")
	}
	var site *Site
	for _, s := range matchedSites {
		if site != nil {
			if len(s.Domain) > len(site.Domain) {
				site = s
			}
		} else {
			site = s
		}
	}

	recordFilters := make([]repos.Filter, 0)

	for i := range domainSplits {
		recordFilters = append(recordFilters, repos.Filter{
			"host": fmt.Sprintf("*.%v", strings.Join(domainSplits[i:], ".")),
		}, repos.Filter{
			"host": strings.Join(domainSplits[i:], "."),
		})
	}

	rec, err := d.recordsRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			"siteId": site.Id,
			"$or":    recordFilters,
		},
		Sort: map[string]interface{}{
			"priority": 1,
		},
	})

	if err != nil {
		return nil, err
	}

	err = d.recordsCache.Set(ctx, host, rec)
	if err != nil {
		fmt.Println(err)
	}

	return rec, nil
}

func (d *domainI) CreateRecord(
	ctx context.Context,
	siteId repos.ID,
	recordType string,
	host string,
	answer string,
	ttl uint32,
	priority int64,
) (*Record, error) {
	create, err := d.recordsRepo.Create(ctx, &Record{
		SiteId:   siteId,
		Type:     recordType,
		Host:     host,
		Answer:   answer,
		TTL:      ttl,
		Priority: priority,
	})
	return create, err
}

func (d *domainI) DeleteRecords(ctx context.Context, host string, siteId string) error {

	d.recordsCache.Drop(ctx, host)

	return d.recordsRepo.DeleteMany(ctx, repos.Filter{
		"host": host,
	})
}

func (d *domainI) AddARecords(ctx context.Context, host string, aRecords []string, siteId string) error {
	var err error

	// fmt.Println(aRecords, host, siteId)
	d.recordsCache.Drop(ctx, host)

	for _, aRecord := range aRecords {
		_, err = d.recordsRepo.Create(ctx, &Record{
			SiteId:   repos.ID(siteId),
			Type:     "A",
			Host:     host,
			Answer:   aRecord,
			TTL:      30,
			Priority: 0,
		})

	}
	return err
}

func fxDomain(
	recordsRepo repos.DbRepo[*Record],
	sitesRepo repos.DbRepo[*Site],
	accountDNSRepo repos.DbRepo[*AccountCName],
	recordsCache cache.Repo[[]*Record],
) Domain {
	return &domainI{
		recordsRepo,
		sitesRepo,
		recordsCache,
		accountDNSRepo,
	}
}

var Module = fx.Module(
	"domain",
	fx.Provide(fxDomain),
)
