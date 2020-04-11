package model

import "time"

type Order struct {
	Id int
	OrderNumber string
	UserId int
	P Payment
	Status OrderStatus
	CreatedTime time.Time
	UpdatedTime time.Time

}

type OrderStatus string

const (
	ORDER_STATUS_UNPAID OrderStatus ="unpaid"
	ORDER_STATUS_PAID = "paid"
	ORDER_STATUS_SHIPPED = "shipped"
	ORDER_STATUS_RECEIVED = "received"
)
