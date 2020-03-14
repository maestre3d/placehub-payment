package rdbms

import (
	"context"
	"database/sql"
	"github.com/maestre3d/placehub-payment/internal/payment/domain"
	"go.uber.org/zap"
)

type PaymentRepository struct {
	logger *zap.SugaredLogger
	conn   *sql.Conn
}

func NewPaymentRepository(logger *zap.SugaredLogger, conn *sql.Conn) *PaymentRepository {
	return &PaymentRepository{logger, conn}
}

func (p *PaymentRepository) FetchAllPayments() ([]*domain.Payment, error) {
	statement := `SELECT * FROM PAYMENTS`
	rows, err := p.conn.QueryContext(context.Background(), statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer p.conn.Close()

	payments := make([]*domain.Payment, 0)
	for rows.Next() {
		payment := new(domain.Payment)
		err := rows.Scan(&payment.ID, &payment.Username, &payment.PlaceID)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
