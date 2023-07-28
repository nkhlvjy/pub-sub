package receiver2

import (
	"log"
	"pub-sub-model/brokers/rabbitmq"
)

type IReceiver2 interface {
	Subscribe(exName string) error
	Receive() (err error)
}

type receiver2 struct {
	sub rabbitmq.ISubscribe
}

func NewReceiver2(name string) IReceiver2 {
	return &receiver2{}
}

func (r2 *receiver2) Subscribe(exName string) error {
	sub, err := rabbitmq.Subscribe("second subscriber", exName)
	r2.sub = sub
	return err
}

func (r2 *receiver2) Receive() (err error) {
	receiverChannel, err := r2.sub.ReceiveMessage()
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
