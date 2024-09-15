package app

import (
	"context"
	"encoding/json"

	"github.com/kloudlite/api/apps/finance/internal/domain"
	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
)

type (
	ChargesConsumer messaging.Consumer
)

func processCharges(ctx context.Context, d domain.Domain, consumer ChargesConsumer, logr logging.Logger) error {
	err := consumer.Consume(func(msg *msgTypes.ConsumeMsg) error {
		logger := logr.WithName("charges")
		logger.Infof("started processing")
		defer func() {
			logger.Infof("finished processing")
		}()
		var charge entities.Charge
		if err := json.Unmarshal(msg.Payload, &charge); err != nil {
			logger.Errorf(err, "could not unmarshal into *Charge")
			return errors.NewE(err)
		}
		logger = logger.WithKV("charge", charge.Id)
		logger.Infof("charge: %s", charge.Id)

		if _, err := d.CreateCharge(ctx, &charge); err != nil {
			logger.Errorf(err, "could not create charge")
			return errors.NewE(err)
		}
		return nil
	}, msgTypes.ConsumeOpts{
		OnError: func(err error) error {
			logr.Errorf(err, "error while consuming message")
			return nil
		},
	})
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
