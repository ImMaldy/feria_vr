package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://tzoeibjf:BxXQ3VAX_Nyl1hW1D9A28hmrnquflDaG@jackal.rmq.cloudamqp.com/tzoeibjf")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("server_1", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	for {
		err := ch.Publish("", "server_1", false, false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte("sent at" + time.Now().String()),
			})

		if err != nil {
			break
		}

		time.Sleep(2 * time.Second)
	}
}
