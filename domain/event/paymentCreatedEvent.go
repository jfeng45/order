package event

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/jfeng45/order/domain/model"
	"strconv"
	"time"
)

type PaymentCreatedEvent struct {
	Id int
	SourceAccount int
	TargetAccount int
	Amount float32
	Status model.PaymentStatus
	PaymentMethod model.PaymentMethod
	OrderNumber string
	CreatedTime time.Time
	CompletionTime time.Time
}

func (pc *PaymentCreatedEvent) NewPaymentCreatedDescriptor() *ycq.EventDescriptor {
	aggregateId := strconv.Itoa(pc.Id)
	var version int
	return ycq.NewEventMessage(aggregateId,pc, &version)
}



