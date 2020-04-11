package command

import (
	"github.com/jfeng45/order/domain/model"
	"time"
)

type MakePaymentCommand struct {
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

func NewMakePaymentCommand(p model.Payment) MakePaymentCommand{
	pc := MakePaymentCommand{p.Id,p.SourceAccount,p.TargetAccount, p.Amount,
		p.Status, p.PaymentMethod, p.OrderNumber, p.CreatedTime,
		p.CompletionTime}
	return pc
}

