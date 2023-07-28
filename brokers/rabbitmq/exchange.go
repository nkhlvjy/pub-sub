package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type IExchange interface {
	PublishMessage(msg map[string]interface{}) error
}

type exchange struct {
	name string
}

func CreateExchange(exName string) (ex IExchange, err error) {
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return nil, err
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Failed to close a channel")
		}
	}(ch)

	err = ch.ExchangeDeclare(exName, "fanout", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")
	return &exchange{
		name: exName,
	}, nil
}

func (e *exchange) PublishMessage(msg map[string]interface{}) error {
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return err
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Failed to close a channel")
		}
	}(ch)
	msgByte, err := json.Marshal(msg)
	failOnError(err, "Failed to marshal a message")
	if err != nil {
		return err
	}
	err = ch.Publish(e.name, "", false, false, amqp.Publishing{
		ContentType: "text/json",
		Body:        msgByte,
	})
	failOnError(err, "Failed to publish a message")
	return err
}
