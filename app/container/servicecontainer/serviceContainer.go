package servicecontainer

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/order/app/container"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/applicationservice/dataservice/orderdata/sqldb"
	"github.com/jfeng45/order/domain/usecase/modifyorder"
	"github.com/jfeng45/order/domain/usecase/searchorder"
	"github.com/jfeng45/order/tool/gdbc"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	var value interface{}
	var found bool
	if value, found = sc.Get(container.DATABASE); !found {
		//logger.Log.Debug("find CacheGrpc key=%v \n", key)
		message := "can't find key= " + container.DATABASE + " in container "
		return nil, errors.New(message)
	}
	dt := value.(gdbc.SqlGdbc)
	pds := sqldb.OrderDataSql{dt}

	if value, found = sc.Get(container.MESSAGING_SERVER); !found {
		//logger.Log.Debug("find CacheGrpc key=%v \n", key)
		message := "can't find key=" + container.MESSAGING_SERVER + " in container "
		return nil, errors.New(message)
	}
	ms := value.(gmessaging.MessagingInterface)

	switch code {
		case container.SEARCH_ORDER_USECASE:
			uc := searchorder.SearchOrderUseCase{&pds}
			logger.Log.Debug("found usecase in container for key:", code)
			return &uc, nil
		case container.MODIFY_ORDER_USECASE:
			uc := modifyorder.ModifyOrderUseCase{&pds, ms}
			logger.Log.Debug("found usecase in container for key:", code)
			return &uc, nil
	    default:
	    	message := "can't find key=" + code + " in container "
			return nil, errors.New(message)
		}
	return nil, nil
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}



