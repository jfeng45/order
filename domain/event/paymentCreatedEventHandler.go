package event

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/domain/model"
	"github.com/jfeng45/order/domain/usecase"
)

type PaymentCreatedEventHandler struct {
	Mouc usecase.ModifyOrderUseCaseInterface
}
func(pc PaymentCreatedEventHandler) Handle (message ycq.EventMessage) {
	switch event := message.Event().(type) {

	case *PaymentCreatedEvent:
		status := model.ORDER_STATUS_PAID
		err := pc.Mouc.UpdatePayment(event.OrderNumber, event.Id,status)
		if err != nil {
			logger.Log.Errorf("error in PaymentCreatedEventHandler:", err)
		}
	default:
		logger.Log.Errorf("event type mismatch in PaymentCreatedEventHandler:")
	}
}

