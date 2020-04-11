package dataservice

import "github.com/jfeng45/order/domain/model"

type OrderDataInterface interface {
	Insert(user *model.Order) (resultUser *model.Order, err error)
	Find(id int) (*model.Order, error)
	CreatePayment( orderNumber string, paymentId int,status string ) (int64, error)
}
