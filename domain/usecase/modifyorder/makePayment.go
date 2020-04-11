package modifyorder

import (
	"github.com/pkg/errors"
	"strconv"
)

func (mpu *ModifyOrderUseCase) MakePayment( orderNumber string, paymentId int,status string ) error {
	numberOfUpdate, err := mpu.OrderDataInterface.CreatePayment(orderNumber, paymentId, status)
	if err!= nil {
		return errors.Wrap(err, "")
	}
	if numberOfUpdate != 1 {
		return errors.New("Make payment failed. rows affected is " + strconv.Itoa(int(numberOfUpdate)))
	}
	return nil
}

