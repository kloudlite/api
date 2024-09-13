package domain

import (
	"fmt"
	"github.com/kloudlite/api/common/fields"
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

	link, err := rp.CreatePaymentLink(&PaymentLinkInput{
		Amount:      req.Amount,
		Currency:    CurrencyUSD,
		ReferenceId: string(p.Id),
		Name:        "Abdhesh",
		Description: "KloudLite Payment",
		AccountNo:   ctx.AccountName,
		Email:       "abdhesh@kloudlite.io",
	})

	if err != nil {
		if err := d.paymentRepo.DeleteById(ctx, p.Id); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error creating payment link: %w", err)
	}

	p.Link = link

	one, err := d.paymentRepo.PatchOne(ctx.Context, repos.Filter{
		fields.Id: p.Id,
	}, repos.Document{
		fc.PaymentPaymentLink: p.Link,
	})
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (d *domain) SyncPaymentStatus(ctx UserContext, paymentId repos.ID) (*entities.Payment, error) {
	p, err := d.paymentRepo.FindById(ctx.Context, paymentId)

	if err != nil {
		return nil, err
	}

	if p.Link.Status == entities.PaymentStatusExpired || p.Link.Status == entities.PaymentStatusCancelled || p.Link.Status == entities.PaymentStatusPaid {
		return p, err
	}

	rp := d.newRazorPay()

	pl, err := rp.GetPaymentLink(p.Link.Id)
	if err != nil {
		return nil, err
	}

	one, err := d.paymentRepo.PatchOne(ctx.Context, repos.Filter{
		fields.Id: p.Id,
	}, repos.Document{
		fc.PaymentPaymentLink: pl,
	})
	if err != nil {
		return nil, err
	}
	return one, nil
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
