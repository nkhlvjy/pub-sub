package publisher

import (
	"log"
	"pub-sub-model/brokers/rabbitmq"
)

type Publisher struct {
	exchange rabbitmq.IExchange
}

func NewPublisher(exName string) *Publisher {
	exchange, _ := rabbitmq.CreateExchange(exName)
	return &Publisher{
		exchange: exchange,
	}
}

func (p *Publisher) Publish(message map[string]interface{}) {
	err := p.exchange.PublishMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}
