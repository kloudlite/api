package rp

import (
	"encoding/json"
	"github.com/razorpay/razorpay-go"
	"os"
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

type PaymentLink struct {
	Id          string `json:"id"`
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
	ShortUrl    string `json:"short_url"`
}

type RazorPay interface {
	CreatePaymentLink(in *PaymentInput) (*PaymentLink, error)
	CancelPaymentLink(linkId string) error
	GetPaymentLink(linkId string) (*PaymentLink, error)
}

type razorPayImpl struct {
	rpClient    *razorpay.Client
	callbackUrl string
}

func (r *razorPayImpl) GetPaymentLink(linkId string) (*PaymentLink, error) {
	fetch, err := r.rpClient.PaymentLink.Fetch(linkId, nil, nil)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(fetch)
	if err != nil {
		return nil, err
	}
	p := &PaymentLink{}
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

func (r *razorPayImpl) CreatePaymentLink(in *PaymentInput) (*PaymentLink, error) {
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
	p := &PaymentLink{}
	err = json.Unmarshal(marshal, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func NewRazorPay() RazorPay {
	r := razorPayImpl{}
	rpKey := os.Getenv("RP_KEY")
	rpSecret := os.Getenv("RP_SECRET")
	callbackUrl := os.Getenv("RP_PAYMENT_CALLBACK_URL")
	r.rpClient = razorpay.NewClient(rpKey, rpSecret)
	r.callbackUrl = callbackUrl
	return &r
}
