package modifyorder

import (
	"github.com/jfeng45/order/app/config"
	"github.com/jfeng45/order/applicationservice/dataservice"
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/order/domain/command"
	"github.com/jfeng45/order/domain/model"
	"github.com/pkg/errors"
)

type ModifyOrderUseCase struct {
	OrderDataInterface dataservice.OrderDataInterface
	Mi           gmessaging.MessagingInterface
}

func (mpu *ModifyOrderUseCase) CreateOrder (order *model.Order ) (*model.Order , error) {
	order, err := mpu.OrderDataInterface.Insert(order)
	if err!= nil {
		return nil, errors.Wrap(err, "")

	}
	//Using a new UpdatePayment command to create payment for this order, the command will be sent to payment service
	mpc := command.NewMakePaymentCommand(&order.P)
	err = mpu.Mi.Publish(config.SUBJECT_MAKE_PAYMENT, mpc)
	if err != nil {
			return nil, errors.Wrap(err, "")
	}
	return order, err
}

