package domain

import (
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	fc "github.com/kloudlite/api/apps/finance/internal/entities/field-constants"
	"github.com/kloudlite/api/pkg/repos"
)

const (
	CurrencyUSD = "USD"
)

func (d *domain) GetWallet(ctx UserContext) (*entities.Wallet, error) {
	resp, err := d.walletRepo.FindOne(ctx.Context, repos.Filter{
		fc.WalletTeamId: ctx.AccountName,
	})

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return d.walletRepo.Create(ctx.Context, &entities.Wallet{
			TeamId:    ctx.AccountName,
			Balance:   0,
			Currency:  CurrencyUSD,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	return resp, nil
}

func (d *domain) ListPayments(ctx UserContext) ([]*entities.Payment, error) {
	return d.paymentRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{
			fc.WalletTeamId: ctx.AccountName,
		},
	})
}

func (d *domain) CreatePayment(ctx UserContext, req *entities.Payment) (*entities.Payment, error) {
	req.Link = nil

	p, err := d.paymentRepo.Create(ctx.Context, &entities.Payment{
		Amount:    req.Amount,
		Link:      req.Link,
		TeamId:    ctx.AccountName,
		Currency:  CurrencyUSD,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	rp := d.newRazorPay()

	link, err := rp.CreatePaymentLink(&PaymentInput{
		Amount:      req.Amount,
		Currency:    req.Currency,
		ReferenceId: string(p.Id),
		Name:        fmt.Sprintf("KloudLite Payment %s", p.Id),
		Description: "KloudLite Payment",
		AccountNo:   ctx.AccountName,
	})
	if err != nil {
		return nil, err
	}

	p.Link = link

	finP, err := d.paymentRepo.UpdateById(ctx.Context, p.Id, &entities.Payment{
		Link:      link,
		Amount:    p.Amount,
		TeamId:    p.TeamId,
		Currency:  p.Currency,
		CreatedAt: p.CreatedAt,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return finP, nil
}

func (d *domain) ValidatePayment(ctx UserContext, paymentId repos.ID) (*entities.Payment, error) {
	p, err := d.paymentRepo.FindById(ctx.Context, paymentId)
	if err != nil {
		return nil, err
	}

	rp := d.newRazorPay()

	pl, err := rp.GetPaymentLink(p.Link.Id)
	if err != nil {
		return nil, err
	}

	finP, err := d.paymentRepo.UpdateById(ctx.Context, p.Id, &entities.Payment{
		Link:      pl,
		Amount:    p.Amount,
		TeamId:    p.TeamId,
		Currency:  p.Currency,
		CreatedAt: p.CreatedAt,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return finP, nil
}

func (d *domain) ListCharges(ctx UserContext) ([]*entities.Charge, error) {
	c, err := d.chargeRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{
			fc.WalletTeamId: ctx.AccountName,
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
