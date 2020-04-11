package model

import "time"

type Payment struct {
	Id int
	SourceAccount int
	TargetAccount int
	Amount float32
	Status PaymentStatus
	PaymentMethod PaymentMethod
	OrderNumber string
	CreatedTime time.Time
	CompletionTime time.Time

}
type PaymentStatus string

const (
	PAYMENT_STATUS_UNCOMPLETED PaymentStatus ="uncompleted"
	PAYMENT_STATUS_COMPLETED  = "completed"
	PAYMENT_STATUS_ABNORMAL = "abnormal"
	PAYMENT_STATUS_CANCELED = "canceled"
)

type PaymentMethod string

const (
	PAYMENT_METHOD_WEIXIN PaymentStatus ="weixin"
	PAYMENT_METHOD_ALIPAY = "alipay"
	PAYMENT_METHOD_CREADIT_CARD = "creaditCard"
	PAYMENT_METHOD_CHECK= "check"
	PAYMENT_METHOD_CASH= "cash"
)