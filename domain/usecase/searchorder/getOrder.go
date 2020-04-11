package searchorder

import (
	"github.com/jfeng45/order/applicationservice/dataservice"
	"github.com/jfeng45/order/domain/model"
)

type SearchOrderUseCase struct {
	OrderDataInterface dataservice.OrderDataInterface
	//EventBus              ycq.EventBus
}

func (souc *SearchOrderUseCase) GetOrder(id int) (*model.Order, error) {
	return souc.OrderDataInterface.Find(id)
}

