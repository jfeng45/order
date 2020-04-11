package usecase

import "github.com/jfeng45/order/domain/model"

type SearchOrderUseCaseInterface interface {
	GetOrder(id int) (*model.Order, error)
}

type ModifyOrderUseCaseInterface interface {
	MakePayment(orderNumber string, paymentId int, status string) error
	CreateOrder(order *model.Order) (*model.Order, error)
}
