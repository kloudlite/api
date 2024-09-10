package domain

import (
	"fmt"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) GetWallet(ctx UserContext) (*entities.Wallet, error) {
	return d.walletRepo.FindOne(ctx.Context, repos.Filter{
		"team_id": ctx.AccountName,
	})
}

func (d *domain) GetPayments(ctx UserContext, walletID repos.ID) ([]*entities.Payment, error) {
	return d.paymentRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{
			"wallet_id": walletID,
			"team_id":   ctx.AccountName,
		},
	})
}

func (d *domain) CreatePayment(ctx UserContext, req *entities.Payment) (*entities.Payment, error) {
	p, err := d.paymentRepo.Create(ctx.Context, req)
	if err != nil {
		return nil, err
	}

	// use p.Id as payment id and initiate payment from payment gateway

	return p, nil
}

func (d *domain) validatePayment(ctx UserContext, paymentId repos.ID) error {
	p, err := d.paymentRepo.FindById(ctx.Context, paymentId)
	if err != nil {
		return err
	}
	// validate payment with p.id on payment gateway
	fmt.Println(p)

	return nil
}

func (d *domain) ListCharges(ctx UserContext, walletID repos.ID) ([]*entities.Charge, error) {
	c, err := d.chargeRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{
			"team_id":   ctx.AccountName,
			"wallet_id": walletID,
		},
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

// internal method
func (d *domain) CreateCharge(ctx UserContext, req *entities.Charge) (*entities.Charge, error) {
	return d.chargeRepo.Create(ctx.Context, req)
}
