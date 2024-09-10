package domain

func (d *domain) GetWallet(ctx UserContext) (*entities.Wallet, error) {
	return d.walletRepo.Get(ctx.UserID)
}

func (d *domain) GetPayments(ctx UserContext, walletID repos.ID) ([]*entities.Payment, error) {
	return d.paymentRepo.GetByWallet(ctx.UserID, walletID)
}

func (d *domain) CreatePayment(ctx UserContext, req *entities.Payment) (*entities.Payment, error) {
	if err := d.validatePayment(ctx, req); err != nil {
		return nil, err
	}

	payment, err := d.paymentRepo.Create(ctx.UserID, req)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (d *domain) validatePayment(ctx UserContext, req *entities.Payment) error {
	if req.Amount <= 0 {
		return ErrInvalidAmount
	}
	if req.Currency == "" {
		return ErrInvalidCurrency
	}
	if req.Description == "" {
		return ErrInvalidDescription
	}
	if req.PaymentType == "" {
		return ErrInvalidPaymentType
	}

	// TODO: validate payment type

	return nil
} 

func (d *domain) ListCharges(ctx UserContext) ([]*entities.Charge, error) {
	return d.chargeRepo.GetByUser(ctx.UserID)
}     

func (d *domain) CreateCharge(ctx UserContext, req *entities.Charge) (*entities.Charge, error) {
	if err := d.validateCharge(ctx, req); err != nil {
		return nil, err
	}

	charge, err := d.chargeRepo.Create(ctx.UserID, req)
	if err != nil {
		return nil, err
	}

	return charge, nil
}
