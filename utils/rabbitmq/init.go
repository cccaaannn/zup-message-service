package rabbitmq

import (
	"errors"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var Connection *amqp.Connection
var publisherChannel *amqp.Channel
var _connectionString string
var err error

func Connect(connectionString string) {
	_connectionString = connectionString // store connectionString to use on reconnect
	Connection, err = amqp.Dial(connectionString)

	if err != nil {
		log.Fatal(err)
		panic("[RabbitMQ] Cannot connect to RabbitMQ")
	}

	publisherChannel, err = Connection.Channel()
	if err != nil {
		log.Fatal(err)
		panic("[RabbitMQ] Cannot connect to RabbitMQ channel")
	}

	log.Printf("[RabbitMQ] Connected to RabbitMQ with %s", connectionString)
}

func Reconnect(retries int) bool {
	if _connectionString == "" {
		log.Println("[RabbitMQ] connectionString is not initialized, 'Connect()' must be called before using 'Reconnect()'")
		return false
	}

	for i := 0; i < retries; i++ {
		log.Println("[RabbitMQ] Attempting to reconnect to RabbitMQ")
		Connect(_connectionString)
		_, err := Connection.Channel()
		if err == nil {
			log.Printf("[RabbitMQ] Reconnected successfully on attempt number (%d) \n", i+1)
			return true
		}

		// This will hang all incoming messages since it is not async, but without RabbitMQ service is dead anyways so it worths the shot
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("[RabbitMQ] Could not been reconnected on attempt number (%d) \n", retries)
	return false
}

func GetChannel() (*amqp.Channel, error) {
	if Connection == nil {
		return nil, errors.New("[RabbitMQ] Connection is nil, most likely never initialized")
	}
	return Connection.Channel()
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
		panic("[RabbitMQ] Cannot create RabbitMQ queue")
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
