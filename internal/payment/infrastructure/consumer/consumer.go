package consumer

import (
	"context"
	"encoding/json"

	"github.com/looksaw/go-orderv2/common/broker"
	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/payment/app"
	"github.com/looksaw/go-orderv2/payment/app/command"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct{
	app app.Application
}

func NewConsumer(app app.Application) *Consumer {
	return &Consumer{
		app: app,
	}
}

func (c *Consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.EventOrderCreate, true, false, false, false, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		logrus.Warnf("fail to consume: queue=%s, err=%v", q.Name, err)
	}

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			c.handleMessage(msg, q, ch)
		}
	}()
	<-forever
}

func (c *Consumer) handleMessage(msg amqp.Delivery, q amqp.Queue, ch *amqp.Channel) {
	logrus.Infof("payment receive a message from %s ,msg=%v", q.Name, string(msg.Body))
	o := &orderpb.Order{}
	if err := json.Unmarshal(msg.Body,o); err != nil {
		_ = msg.Nack(false,false)
		return 
	}
	if _ , err := c.app.Command.CreatePayment.Handle(context.TODO(),command.CreatePayment{Order: o}); err != nil {
		logrus.Infof("failed to create order , err= %v",err)
		_ = msg.Nack(false,false)
		return 
	}
	_ = msg.Ack(false)
}
