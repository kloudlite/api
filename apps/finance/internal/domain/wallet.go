package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	fc "github.com/kloudlite/api/apps/finance/internal/entities/field-constants"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/accounts"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/repos"
)

const (
	CurrencyUSD = "USD"
)

func (d *domain) canPerformActionInAccount(ctx UserContext, action iamT.Action) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(action),
	})
	if err != nil {
		return errors.NewE(err)
	}
	if !co.Status {
		return errors.Newf("unauthorized to perform action %q in account %q", action, ctx.AccountName)
	}
	return nil
}

func (d *domain) GetWallet(ctx UserContext) (*entities.Wallet, error) {

	if err := d.canPerformActionInAccount(ctx, iamT.GetAccount); err != nil {
		return nil, err
	}

	resp, err := d.walletRepo.FindOne(ctx.Context, repos.Filter{
		fc.WalletTeamId: ctx.AccountName,
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if resp == nil {
		return d.walletRepo.Create(ctx.Context, &entities.Wallet{
			TeamId:   ctx.AccountName,
			Balance:  0,
			Currency: CurrencyUSD,
		})
	}

	return resp, nil
}

func (d *domain) ListPayments(ctx UserContext) ([]*entities.Payment, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetAccount); err != nil {
		return nil, err
	}

	return d.paymentRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{
			fc.WalletTeamId: ctx.AccountName,
		},
	})
}

func (d *domain) CreatePayment(ctx UserContext, req *entities.Payment) (*entities.Payment, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.DeleteAccount); err != nil {
		return nil, err
	}

	if req.Amount <= 0 {
		return nil, errors.Newf("invalid amount %f", req.Amount)
	}

	gao, err := d.accountsCli.GetAccount(ctx, &accounts.GetAccountIn{
		UserId:      string(ctx.UserId),
		AccountName: ctx.AccountName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	// account email address supposed to come from the gao
	fmt.Println(gao)

	p, err := d.paymentRepo.Create(ctx.Context, &entities.Payment{
		Amount:    req.Amount,
		Link:      nil,
		TeamId:    ctx.AccountName,
		Currency:  CurrencyUSD,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	rp := d.newRazorPay()
	link, err := rp.CreatePaymentLink(&PaymentLinkInput{
		Amount:      req.Amount,
		Currency:    CurrencyUSD,
		ReferenceId: string(p.Id),
		Name:        ctx.UserName,
		Description: "KloudLite Payment",
		AccountNo:   ctx.AccountName,
		Email:       ctx.UserEmail,
	})

	if err != nil {
		if err := d.paymentRepo.DeleteById(ctx, p.Id); err != nil {
			return nil, errors.NewE(err)
		}

		return nil, fmt.Errorf("error creating payment link: %w", err)
	}

	one, err := d.paymentRepo.PatchOne(ctx.Context, repos.Filter{
		fields.Id: p.Id,
	}, repos.Document{
		fc.PaymentPaymentLink: link,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	return one, nil
}

func (d *domain) SyncPaymentStatus(ctx UserContext, paymentId repos.ID) (*entities.Payment, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.DeleteAccount); err != nil {
		return nil, errors.NewE(err)
	}

	p, err := d.paymentRepo.FindById(ctx.Context, paymentId)

	if err != nil {
		return nil, errors.NewE(err)
	}

	if p.Link.Status == entities.PaymentStatusExpired || p.Link.Status == entities.PaymentStatusCancelled || p.Link.Status == entities.PaymentStatusPaid {
		return p, err
	}

	rp := d.newRazorPay()

	pl, err := rp.GetPaymentLink(p.Link.Id)
	if err != nil {
		return nil, errors.NewE(err)
	}

	// handle this block carefully
	if pl.Status == entities.PaymentStatusPaid && p.Link.Status != entities.PaymentStatusPaid {

		wl, err := d.walletRepo.FindOne(ctx, repos.Filter{
			fc.WalletTeamId: ctx.AccountName,
		})

		if err != nil {
			return nil, errors.NewE(err)
		}

		if _, err = d.walletRepo.PatchOne(ctx, repos.Filter{
			fields.Id:       wl.Id,
			fc.WalletTeamId: ctx.AccountName,
		}, repos.Document{
			fc.WalletBalance: wl.Balance + p.Amount,
		}); err != nil {
			return nil, errors.NewE(err)
		}
	}

	/*
		potential risk
		- if the wallet is updated but link is not updated then we will have issue with balance
	*/

	one, err := d.paymentRepo.PatchOne(ctx.Context, repos.Filter{
		fields.Id: p.Id,
	}, repos.Document{
		fc.PaymentPaymentLink: pl,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	return one, nil
}

func (d *domain) ListCharges(ctx UserContext) ([]*entities.Charge, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetAccount); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.chargeRepo.Find(ctx.Context, repos.Query{
		Filter: repos.Filter{fc.WalletTeamId: ctx.AccountName},
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

// internal method
func (d *domain) CreateCharge(ctx context.Context, req *entities.Charge) (*entities.Charge, error) {
	return d.chargeRepo.Create(ctx, req)
}
