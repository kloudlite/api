package domain

import (
	"encoding/json"
	"fmt"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/razorpay/razorpay-go"
)

type PaymentLinkInput struct {
	Name        string
	Email       string
	AccountNo   string
	ReferenceId string
	Description string
	Amount      int
	Currency    string
}

type RazorPay interface {
	CreatePaymentLink(in *PaymentLinkInput) (*entities.PaymentLink, error)
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

func (r *razorPayImpl) CreatePaymentLink(in *PaymentLinkInput) (*entities.PaymentLink, error) {
	data := map[string]any{
		"amount":       in.Amount,
		"currency":     in.Currency,
		"reference_id": in.ReferenceId,
		"description":  in.Description,
		"customer": map[string]any{
			"name":  in.Name,
			"email": in.Email,
		},
		"notify": map[string]any{
			"email": true,
		},
		"reminder_enable": true,
		"notes": map[string]any{
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
	err = json.Unmarshal(marshal, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (d *domain) newRazorPay() RazorPay {

	r := razorPayImpl{}
	rpKey := d.Env.RP_KEY
	rpSecret := d.Env.RP_SECRET
	callbackUrl := d.Env.RP_PAYMENT_CALLBACK_URL
	r.rpClient = razorpay.NewClient(rpKey, rpSecret)
	r.callbackUrl = callbackUrl

	fmt.Printf("%s, %s, %s\n", rpKey, rpSecret, callbackUrl)
	return &r
}

