package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maestre3d/placehub-payment/internal/payment/application"
	"github.com/maestre3d/placehub-payment/internal/shared/domain/util"
	"go.uber.org/zap"
	"net/http"
)

type PaymentHandler struct {
	mux     *gin.RouterGroup
	logger  *zap.SugaredLogger
	payment *application.PaymentUseCase
}

func InitPaymentHandler(logger *zap.SugaredLogger, mux *gin.RouterGroup, payment *application.PaymentUseCase) error {
	logger.Infow("Payment_Handler",
		"status", "started",
	)

	paymentHandler := &PaymentHandler{
		mux:     mux,
		logger:  logger,
		payment: payment,
	}

	paymentHandler.setRoutes()

	return nil
}

func (p *PaymentHandler) setRoutes() *gin.RouterGroup {
	payments := p.mux.Group("/payment")
	{
		payments.GET("/", p.GetPayments)
		payments.GET("/:id", p.GetPayment)
	}

	return payments
}

func (p *PaymentHandler) GetPayment(c *gin.Context) {
	c.JSON(http.StatusOK, &util.Response{"Hello there from Gin Payment"})
}

func (p *PaymentHandler) GetPayments(c *gin.Context) {
	payments, err := p.payment.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &util.Response{err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
