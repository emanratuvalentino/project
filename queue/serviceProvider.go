package queue

import (
	"code-be-docudigital/config"
	"code-be-docudigital/helper"
	"fmt"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel
var err error

func Init() {
	// Connect to RabbitMQ server
	// fmt.Println("Connecting to RabbitMQ ...")
	configuration := config.GetConfig()
	connectString := fmt.Sprintf("amqp://%s:%s@%s:%s/", configuration.AMQP_USERNAME, configuration.AMQP_PASSWORD, configuration.AMQP_HOST, configuration.AMQP_PORT)

	conn, err := amqp.Dial(connectString) //Insert the  connection string
	helper.FailOnError(err, "RabbitMQ connection failure", "RabbitMQ Connection Established")
	// defer conn.Close()

	//Connect to the channel
	ch, err = conn.Channel()
	helper.FailOnError(err, "Failed to open a channel", "Opened the channel")
	// defer ch.Close()

}

func Manager() *amqp.Channel {
	return ch
}
