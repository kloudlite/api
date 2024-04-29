package app

import (
	"context"
	"embed"
	"fmt"
	"text/template"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofiber/fiber/v2"
	"github.com/kloudlite/api/constants"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"github.com/kloudlite/api/pkg/kv"

	"github.com/kloudlite/api/apps/comms/internal/app/graph"
	"github.com/kloudlite/api/apps/comms/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/comms/internal/domain"

	"github.com/kloudlite/api/apps/comms/internal/domain/entities"
	"github.com/kloudlite/api/apps/comms/internal/env"
	"github.com/kloudlite/api/apps/comms/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/messaging"
	msg_nats "github.com/kloudlite/api/pkg/messaging/nats"
	"github.com/kloudlite/api/pkg/nats"
	"github.com/kloudlite/api/pkg/repos"

	"github.com/kloudlite/api/pkg/grpc"

	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/comms"
	"go.uber.org/fx"
)

type NotificationConsumer messaging.Consumer

type CommsGrpcServer grpc.Server

type EmailTemplatesDir struct {
	embed.FS
}

type EmailTemplate struct {
	Subject   string
	Html      *template.Template
	PlainText *template.Template
}

type AccountInviteEmail *EmailTemplate
type ProjectInviteEmail *EmailTemplate
type RestPasswordEmail *EmailTemplate
type UserVerificationEmail *EmailTemplate
type WelcomeEmail *EmailTemplate

func parseTemplate(et EmailTemplatesDir, templateName string, subject string) (*EmailTemplate, error) {
	txtFile, err := et.ReadFile(fmt.Sprintf("email-templates/%v/email.txt", templateName))
	if err != nil {
		return nil, errors.NewE(err)
	}
	txt, err := template.New("email-text").Parse(string(txtFile))
	if err != nil {
		return nil, errors.NewE(err)
	}

	htmlFile, err := et.ReadFile(fmt.Sprintf("email-templates/%v/email.html", templateName))
	if err != nil {
		return nil, errors.NewE(err)
	}
	html, err := template.New(templateName).Parse(string(htmlFile))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return &EmailTemplate{
		Subject:   subject,
		Html:      html,
		PlainText: txt,
	}, nil
}

var Module = fx.Module("app",
	repos.NewFxMongoRepo[*entities.NotificationConf]("nconfs", "prj", entities.NotificationConfIndexes),
	repos.NewFxMongoRepo[*entities.Subscription]("subscriptions", "prj", entities.SubscriptionIndexes),
	repos.NewFxMongoRepo[*types.Notification]("notifications", "prj", entities.SubscriptionIndexes),

	fx.Provide(func(jc *nats.JetstreamClient, ev *env.Env, logger logging.Logger) (NotificationConsumer, error) {
		topic := string(common.NotificationTopicName)
		consumerName := "ntfy:message"
		return msg_nats.NewJetstreamConsumer(context.TODO(), jc, msg_nats.JetstreamConsumerArgs{
			Stream: ev.NotificationNatsStream,
			ConsumerConfig: msg_nats.ConsumerConfig{
				Name:        consumerName,
				Durable:     consumerName,
				Description: "this consumer reads message from a subject dedicated to errors, that occurred when the resource was applied at the agent",
				FilterSubjects: []string{
					topic,
				},
			},
		})
	}),

	fx.Provide(func(et EmailTemplatesDir) (AccountInviteEmail, error) {
		return parseTemplate(et, "account-invite", "[Kloudlite] Account Invite")
	}),
	fx.Provide(func(et EmailTemplatesDir) (ProjectInviteEmail, error) {
		return parseTemplate(et, "project-invite", "[Kloudlite] Project Invite")
	}),
	fx.Provide(func(et EmailTemplatesDir) (RestPasswordEmail, error) {
		return parseTemplate(et, "reset-password", "[Kloudlite] Reset Password")
	}),
	fx.Provide(func(et EmailTemplatesDir) (UserVerificationEmail, error) {
		return parseTemplate(et, "user-verification", "[Kloudlite] Verify Email")
	}),
	fx.Provide(func(et EmailTemplatesDir) (WelcomeEmail, error) {
		return parseTemplate(et, "welcome", "[Kloudlite] Welcome to Kloudlite")
	}),

	fx.Provide(newCommsSvc),

	fx.Invoke(func(server CommsGrpcServer, commsServer comms.CommsServer) {
		comms.RegisterCommsServer(server, commsServer)
	}),

	fx.Invoke(
		func(server httpServer.Server, d domain.Domain, sessionRepo kv.Repo[*common.AuthSession], ev *env.Env) {
			gqlConfig := generated.Config{Resolvers: &graph.Resolver{Domain: d, Env: ev}}

			gqlConfig.Directives.IsLoggedInAndVerified = func(ctx context.Context, _ interface{}, next graphql.Resolver) (res interface{}, err error) {
				sess := httpServer.GetSession[*common.AuthSession](ctx)
				if sess == nil {
					return nil, fiber.ErrUnauthorized
				}

				if !sess.UserVerified {
					return nil, &fiber.Error{
						Code:    fiber.StatusForbidden,
						Message: "user's email is not verified",
					}
				}

				return next(context.WithValue(ctx, "user-session", sess))
			}

			gqlConfig.Directives.HasAccount = func(ctx context.Context, _ interface{}, next graphql.Resolver) (res interface{}, err error) {
				sess := httpServer.GetSession[*common.AuthSession](ctx)
				if sess == nil {
					return nil, fiber.ErrUnauthorized
				}
				m := httpServer.GetHttpCookies(ctx)
				klAccount := m[ev.AccountCookieName]
				if klAccount == "" {
					return nil, errors.Newf("no cookie named %q present in request", ev.AccountCookieName)
				}

				nctx := context.WithValue(ctx, "user-session", sess)
				nctx = context.WithValue(nctx, "account-name", klAccount)
				return next(nctx)
			}

			schema := generated.NewExecutableSchema(gqlConfig)
			server.SetupGraphqlServer(schema, httpServer.NewReadSessionMiddleware(sessionRepo, constants.CookieName, constants.CacheSessionPrefix))
		},
	),

	fx.Invoke(func(lf fx.Lifecycle, consumer NotificationConsumer, d domain.Domain, logr logging.Logger) {
		lf.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					err := processNotification(ctx, d, consumer, logr)
					if err != nil {
						logr.Errorf(err, "could not process notifications")
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}),
)
