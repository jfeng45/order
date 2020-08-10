package main

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/order/app"
	"github.com/jfeng45/order/app/config"
	"github.com/jfeng45/order/app/container"
	"github.com/jfeng45/order/app/container/containerhelper"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/domain/event"
	"github.com/jfeng45/order/domain/model"
	"github.com/jfeng45/order/tool/timea"
	"log"
	"runtime"
	"time"
)

func main() {
	c, err := app.InitApp()
	if err != nil {
		log.Println("err:", err)
	}
	go testSubscribe(c)
	time.Sleep(1000)
	testMySql(c)
	runtime.Goexit()
}

func testSubscribe(c container.Container) {
	var value interface{}
	var found bool
	if value, found = c.Get(container.MESSAGING_SERVER); !found {
		message := "can't find key= " + container.MESSAGING_SERVER +" in container "
		logger.Log.Errorf(message)
	}
	ms := value.(gmessaging.MessagingEncodedInterface)
	if value, found = c.Get(container.EVENT_BUS); !found {
		message := "can't find key=" + container.EVENT_BUS + " in container "
		logger.Log.Errorf("err:",message)
	}
	eb := value.(ycq.EventBus)
	subject := config.SUBJECT_PAYMENT_CREATED
	_, err := ms.Subscribe(subject, func(pce event.PaymentCreatedEvent) {
		cpm := pce.NewPaymentCreatedDescriptor()
		logger.Log.Debug("payload:",pce)
		eb.PublishEvent(cpm)
	})
	if err != nil {
		logger.Log.Errorf("err:",err)
	}
	log.Printf("Listening on [%s]", subject)
	runtime.Goexit()
}

func testMySql(c container.Container) {
	//testGetOrder(c)
	//testUpdatePayment(c)
	testCreateOrder(c)
}

func testGetOrder(c container.Container) {
	id := 1
	souc, err := containerhelper.BuildSearchOrderUseCase(c)
	if err != nil {
		logger.Log.Fatalf("getOrderUseCase interface build failed:%+v\n", err)
	}
	order, err := souc.GetOrder(id)
	if err != nil {
		logger.Log.Errorf("getOrder failed failed:%+v\n", err)
	}
	logger.Log.Info("find order:", order)

}

func testUpdatePayment(c container.Container) {
	uouc, err := containerhelper.BuildModifyOrderUseCase(c)
	if err != nil {
		logger.Log.Fatalf("UpdateOrderUseCase interface build failed:%+v\n", err)
	}
	err = uouc.UpdatePayment("2", 222, "created2")
	if err != nil {
		logger.Log.Errorf("UpdateOrderUseCase failed failed:%+v\n", err)
	}
	logger.Log.Info("UpdateOrderUseCase :")
}

func testCreateOrder(c container.Container) {
	mouc, err := containerhelper.BuildModifyOrderUseCase(c)
	if err != nil {
		logger.Log.Fatalf("createOrderUseCase interface build failed:%+v\n", err)
	}
	created, err := time.Parse(timea.FORMAT_ISO8601_DATE_TIME, "2020-05-09 15:04:05")
	updated, err := time.Parse(timea.FORMAT_ISO8601_DATE, "2020-02-17 15:04:05")
	var completed time.Time
	orderNumber := "4"
	p := model.Payment{0,2,22,22,
		model.PAYMENT_STATUS_UNCOMPLETED,model.PAYMENT_METHOD_ALIPAY,orderNumber, time.Now(),
		completed}
	var id int
	o := model.Order{id,orderNumber,1,p,
		model.ORDER_STATUS_UNPAID, created, updated}
	order, err := mouc.CreateOrder(&o)
	if err != nil {
		logger.Log.Errorf("createOrderUseCase failed failed:%+v\n", err)
	}
	logger.Log.Info("find order:", order)
}

