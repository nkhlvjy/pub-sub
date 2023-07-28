package rabbitmq

import (
	"log"
	"sync"
)

import (
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var connection *amqp.Connection

func init() {
	once := &sync.Once{}
	createConnection(once)
}

func createConnection(once *sync.Once) {
	once.Do(
		func() {
			amqpURI := "amqp://guest:guest@localhost:5672/"
			conn, err := amqp.Dial(amqpURI)
			failOnError(err, "Failed to connect to RabbitMQ")
			if err != nil {
				return
			}

			go func(conn *amqp.Connection) {
				select {
				case err := <-conn.NotifyClose(make(chan *amqp.Error)):
					log.Println(err)
					panic(err)
				}
			}(conn)
			connection = conn
		})

}
