package queue

import (
	"code-be-docudigital/helper"
	"fmt"
)

func ReceivedMessage() {
	ch := Manager()

	q, err := ch.QueueDeclare(
		"DemoQueue", //name
		//"ha.monitoring",
		true,
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //arguements
	)
	helper.FailOnError(err, "Failed to declare the queue", "Declared the queue")

	msgs, err := ch.Consume(
		q.Name, //queue
		"",     //consumer
		true,   //auto-ack
		false,  //exclusive
		false,  //no-local
		false,  //no-wait
		nil,    //args
	)
	helper.FailOnError(err, "Failed to register a consumer ", "Registered the consumer")

	msgCount := 0
	go func() {
		for d := range msgs {

			msgCount++

			fmt.Printf("\nMessage Count: %d, Message Body: %s\n", msgCount, d.Body)

		}
	}()
}
