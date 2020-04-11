package event

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/domain/model"
	"github.com/jfeng45/order/domain/usecase"
)

type PaymentCreatedEventHandler struct {
	Uouc usecase.ModifyOrderUseCaseInterface
}
func(pc PaymentCreatedEventHandler) Handle (message ycq.EventMessage) {
	switch event := message.Event().(type) {

	case *PaymentCreatedEvent:
		status := model.PAYMENT_STATUS_COMPLETED
		err := pc.Uouc.MakePayment(event.OrderNumber, event.Id,status)
		if err != nil {
			logger.Log.Errorf("error in PaymentCreatedEventHandler:", err)
		}
	default:
		logger.Log.Errorf("event type mismatch in PaymentCreatedEventHandler:")
	}
}

