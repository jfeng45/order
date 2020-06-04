package app

import (
	"database/sql"
	ycq "github.com/jetbasrawi/go.cqrs"
	logConfig "github.com/jfeng45/glogger/config"
	logFactory "github.com/jfeng45/glogger/factory"
	"github.com/jfeng45/gmessaging"
	gmessagingConfig "github.com/jfeng45/gmessaging/config"
	gmessagingFactory "github.com/jfeng45/gmessaging/factory"
	"github.com/jfeng45/order/app/config"
	"github.com/jfeng45/order/app/container"
	"github.com/jfeng45/order/app/container/containerhelper"
	"github.com/jfeng45/order/app/container/servicecontainer"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/domain/event"
	"github.com/jfeng45/order/tool/gdbc"
	"github.com/jfeng45/order/tool/gdbc/databasehandler"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

// InitApp initialize the application container and load resources like JDBC, Messaging server and so on.
// InitApp only needs to be called once. If the configuration changes, you can call it again to reinitialize the app.
// "filename" is optional, it is provided, then the configuration is loaded from the file and saved them in appConfig,
// otherwise it is provided in code.
// You can also hard code the configuration values in code.
func InitApp(filename...string) (container.Container, error) {
	err := initLogger()
	if err != nil {
		return nil, err
	}
	return initContainer()
}

func initContainer() (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{factoryMap}
	gdbc, err :=initGdbc()
	if err != nil {
		return nil,err
	}
	c.Put(container.DATABASE, gdbc)
	ec, err := initMessagingService()
	if err != nil {
		return nil, err
	}
	c.Put(container.MESSAGING_SERVER, ec)
	eb := initEventBus()
	c.Put(container.EVENT_BUS, eb)
	loadEventHandler(&c)
	return &c, nil
}

func initLogger () error{
	lc := logConfig.Logging{logConfig.ZAP, logConfig.DEBUG, config.LOG_ENABLE_CALLER}
	log, err := logFactory.Build(&lc)
	if err != nil {
		return errors.Wrap(err, "loadLogger")
	}
	logger.SetLogger(log)
	return nil
}

func initGdbc() (gdbc.SqlGdbc,error) {

	db, err := sql.Open(config.DB_DRIVER_NAME, config.DB_SOURCE_NAME)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt := databasehandler.SqlDBTx{DB: db}
	return &dt, nil
}

func initEventBus() ycq.EventBus {
	// Create the EventBus
	eventBus := ycq.NewInternalEventBus()
	return eventBus
}

func loadEventHandler(c container.Container) error {
	var value interface{}
	var found bool

	rluf, err := containerhelper.BuildModifyOrderUseCase(c)
	if err != nil {
		return err
	}
	pceh := event.PaymentCreatedEventHandler{rluf}
	if value, found = c.Get(container.EVENT_BUS); !found {
		//logger.Log.Debug("find CacheGrpc key=%v \n", key)
		message := "can't find key=" + container.EVENT_BUS + " in container "
		return errors.New(message)
	}
	eb := value.(ycq.EventBus)
	eb.AddHandler(&pceh,&event.PaymentCreatedEvent{})
	return nil
}
//func initMessagingService() (gmessaging.MessagingInterface, error) {
//	url := config.MESSAGING_SERVER_URL
//	nc, err :=nats.Connect(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	//defer nc.Close()
//	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
//	if err != nil {
//		return nil, err
//	}
//	n := nat.Nat{ec}
//	return &n, nil
//	//defer ec.Close()
//}

func initMessagingService() (gmessaging.MessagingInterface, error) {
	config := gmessagingConfig.Messaging{gmessagingConfig.NATS_ENCODED, config.MESSAGING_SERVER_URL, nats.JSON_ENCODER}
	return gmessagingFactory.Build(&config)
}
