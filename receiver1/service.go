package receiver1

import (
	"log"
	"pub-sub-model/brokers/rabbitmq"
)

type IReceiver1 interface {
	Subscribe(exName string) error
	Receive() (err error)
}

type receiver1 struct {
	sub rabbitmq.ISubscribe
}

func NewReceiver1(name string) IReceiver1 {
	return &receiver1{}
}

func (r1 *receiver1) Subscribe(exName string) error {
	sub, err := rabbitmq.Subscribe("first subscriber", exName)
	r1.sub = sub
	return err
}

func (r1 *receiver1) Receive() (err error) {
	receiverChannel, err := r1.sub.ReceiveMessage()
	if err != nil {
		return
	}
	go func() {
		for msg := range receiverChannel {
			log.Println(string(msg.Body))
		}
	}()
	return
}
