package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var Connection *amqp.Connection
var publisherChannel *amqp.Channel
var err error

func Connect(connectionString string) {
	Connection, err = amqp.Dial(connectionString)

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to RabbitMQ")
	}

	publisherChannel, err = Connection.Channel()
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to RabbitMQ channel")
	}

	log.Printf("Connected to RabbitMQ with %s", connectionString)
}

func CreateQueue(channel *amqp.Channel, queueName string) {
	_, err = channel.QueueDeclare(
		queueName, // queue name
		false,     // durable
		true,      // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err)
		panic("Cannot create RabbitMQ queue")
	}
}

func Disconnect() {
	publisherChannel.Close()
	Connection.Close()
}

func PublishMessage(messageBody []byte, queueName string) {
	CreateQueue(publisherChannel, queueName)

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(messageBody),
	}

	// Attempt to publish a message to the queue.
	if err := publisherChannel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		message,   // message to publish
	); err != nil {
		log.Println(err)
	}
}

func GetConsumer(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	CreateQueue(channel, queueName)

	// Subscribing to QueueService1 for getting messages.
	messagesConsumer, err := channel.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Println(err)
	}
	return messagesConsumer, err
}
