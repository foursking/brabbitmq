package brabbitmq

import (
	"encoding/json"
	"fmt"
	"time"
	"log"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// Exchange type
var (
	ExchangeDirect  = amqp.ExchangeDirect
	ExchangeFanout  = amqp.ExchangeFanout
	ExchangeTopic   = amqp.ExchangeTopic
	ExchangeHeaders = amqp.ExchangeHeaders
)

// DeliveryMode
var (
	Transient  uint8 = amqp.Transient
	Persistent uint8 = amqp.Persistent
)

var (
	/*
	url = fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		setting.Config().RabbitMq.User,
		setting.Config().RabbitMq.Password,
		setting.Config().RabbitMq.Host,
		setting.Config().RabbitMq.Port,
		setting.Config().RabbitMq.Vhost)
		*/
	exchangeName                           = "service"
	offLineReconnectInterval time.Duration = 10
	retryTimes                             = 5
)

func dealError(err error) error {
	if err != nil {
		log.Pringln(err.Error())
	}
	return err
}

func MqPack(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	return bytes, err
}

func MqUnpack(bytes []byte, data interface{}) error {
	err := json.Unmarshal(bytes, data)
	return err
}
