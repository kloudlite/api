package domain

import (
	"encoding/json"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/razorpay/razorpay-go"
)

type PaymentInput struct {
	Name        string
	Email       string
	AccountNo   string
	ReferenceId string
	Description string
	Amount      int
	Currency    string
}


type RazorPay interface {
	CreatePaymentLink(in *PaymentInput) (*entities.PaymentLink, error)
	CancelPaymentLink(linkId string) error
	GetPaymentLink(linkId string) (*entities.PaymentLink, error)
}

type razorPayImpl struct {
	rpClient    *razorpay.Client
	callbackUrl string
}

func (r *razorPayImpl) GetPaymentLink(linkId string) (*entities.PaymentLink, error) {
	fetch, err := r.rpClient.PaymentLink.Fetch(linkId, nil, nil)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(fetch)
	if err != nil {
		return nil, err
	}
	p := &entities.PaymentLink{}
	err = json.Unmarshal(marshal, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *razorPayImpl) CancelPaymentLink(linkId string) error {
	_, err := r.rpClient.PaymentLink.Cancel(linkId, nil, nil)
	return err
}

func (r *razorPayImpl) CreatePaymentLink(in *PaymentInput) (*entities.PaymentLink, error) {
	data := map[string]interface{}{
		"amount":       in.Amount,
		"currency":     in.Currency,
		"reference_id": in.ReferenceId,
		"description":  in.Description,
		"customer": map[string]interface{}{
			"name":  in.Name,
			"email": in.Email,
		},
		"notify": map[string]interface{}{
			"email": true,
		},
		"reminder_enable": true,
		"notes": map[string]interface{}{
			"account": in.AccountNo,
		},
		"callback_url":    r.callbackUrl,
		"callback_method": "get",
	}

	body, err := r.rpClient.PaymentLink.Create(data, nil)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(body)
	p := &entities.PaymentLink{}
	if err = json.Unmarshal(marshal, p); err != nil {
		return nil, err
	}

	return p, nil
}

func (d *domain) newRazorPay() RazorPay {
	r := razorPayImpl{
		rpClient:    razorpay.NewClient(d.Env.RP_KEY, d.Env.RP_SECRET),
		callbackUrl: d.Env.RP_PAYMENT_CALLBACK_URL,
	}
	return &r
}
