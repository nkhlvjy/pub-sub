package main

import (
	publisher2 "pub-sub-model/publisher"
	"pub-sub-model/receiver1"
	"pub-sub-model/receiver2"
	"time"
)

func main() {

	publisher := publisher2.NewPublisher("test")
	publisher.Publish(map[string]interface{}{"test": "test"})
	r1 := receiver1.NewReceiver1("receiver1")
	r2 := receiver2.NewReceiver2("receiver2")

	err := r1.Subscribe("test")
	if err != nil {
		return
	}
	err = r2.Subscribe("test")
	if err != nil {
		return
	}
	err = r1.Receive()
	if err != nil {
		return
	}
	err = r2.Receive()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
}
