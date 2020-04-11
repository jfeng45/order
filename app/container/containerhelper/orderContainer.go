package containerhelper

import (
	"github.com/jfeng45/order/app/container"
	"github.com/jfeng45/order/domain/usecase"
	"github.com/pkg/errors"
)

func BuildModifyOrderUseCase(c container.Container) (usecase.ModifyOrderUseCaseInterface, error) {
	key := container.MODIFY_ORDER_USECASE
	value, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.ModifyOrderUseCaseInterface), nil
}

func BuildSearchOrderUseCase(c container.Container) (usecase.SearchOrderUseCaseInterface, error) {
	key := container.SEARCH_ORDER_USECASE
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.SearchOrderUseCaseInterface), nil
}

