package application

import (
	"github.com/maestre3d/placehub-payment/internal/payment/domain"
	"github.com/maestre3d/placehub-payment/internal/payment/infrastructure/persistence/rdbms"
)

type PaymentUseCase struct {
	repository *rdbms.PaymentRepository
}

func NewPaymentUseCase(repository *rdbms.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{repository}
}

func (p *PaymentUseCase) GetAllPayments() ([]*domain.Payment, error) {
	return p.repository.FetchAllPayments()
}
