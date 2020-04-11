package config

import "github.com/nats-io/nats.go"

const (
	LOG_CODE             string = "zap"
	LOG_LEVEL            string = "debug"
	LOG_ENABLE_CALLER    bool   = true

	DB_DRIVER_NAME       string = "mysql"
	DB_SOURCE_NAME       string ="root:@tcp(localhost:4333)/fasp?charset=utf8"

	MESSAGING_SERVER_URL string = nats.DefaultURL

	SUBJECT_PAYMENT_CREATED string ="payment.paymentCreated"
	SUBJECT_MAKE_PAYMENT    string ="payment.makePayment"
)