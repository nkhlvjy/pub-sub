package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type ISubscribe interface {
	ReceiveMessage() (delivery <-chan amqp.Delivery, err error)
}

type subscriber struct {
	queue string
}

func Subscribe(queueName, exchange string) (iSub ISubscribe, err error) {
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Failed to close a channel")
		}
	}(ch)
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	if err != nil {
		return
	}
	err = ch.QueueBind(queueName, "", exchange, false, nil)
	failOnError(err, "Failed to bind a queue")
	return &subscriber{
		queue: queueName,
	}, err
}

func (sub *subscriber) ReceiveMessage() (delivery <-chan amqp.Delivery, err error) {
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return
	}

	delivery, err = ch.Consume(sub.queue, "", true, false, false, false, nil)
	failOnError(err, "Failed to consume from opened channel")
	if err != nil {
		return
	}
	go func() {
		select {
		case err := <-ch.NotifyClose(make(chan *amqp.Error)):
			log.Println(err)
			panic(err)
		}
	}()
	return delivery, err
}
