package rabbitmq

import (
	"errors"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var publisherChannel *amqp.Channel
var _connectionString string
var retry_count = 5

func Connect(connectionString string) {
	_connectionString = connectionString // store connectionString to use on reconnect
	var err error = nil
	connection, err = amqp.Dial(connectionString)

	if err != nil {
		log.Fatal(err)
		panic("[RabbitMQ] Cannot connect to RabbitMQ")
	}

	publisherChannel, err = connection.Channel()
	if err != nil {
		log.Println("[RabbitMQ] Cannot connect to RabbitMQ channel. error: ", err)
	}

	log.Printf("[RabbitMQ] Connected to RabbitMQ with %s", connectionString)
}

func Disconnect() {
	publisherChannel.Close()
	connection.Close()
}

func Reconnect(retries int) bool {
	if _connectionString == "" {
		log.Println("[RabbitMQ] connectionString is not initialized, 'Connect()' must be called before using 'Reconnect()'")
		return false
	}

	for i := 0; i < retries; i++ {
		log.Println("[RabbitMQ] Attempting to reconnect to RabbitMQ")
		Connect(_connectionString)
		_, err := connection.Channel()
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

/*
 * This function has retry logic in it since all communication is done by first retrieving channel from this function .
 */
func GetChannel() (*amqp.Channel, error) {
	// Retry on nil connection
	if connection == nil {
		reconnect_result := Reconnect(retry_count)
		if !reconnect_result {
			return nil, errors.New("[RabbitMQ] could not reconnected")
		}
	}

	// Retry on channel error
	_, err := connection.Channel()
	if err != nil {
		reconnect_result := Reconnect(retry_count)
		if !reconnect_result {
			return nil, errors.New("[RabbitMQ] could not reconnected")
		}
	}

	return connection.Channel()
}

func createQueue(channel *amqp.Channel, queueName string) {
	_, err := channel.QueueDeclare(
		queueName, // queue name
		false,     // durable
		true,      // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("[RabbitMQ] Cannot create RabbitMQ queue. error: ", err)
	}
}

func PublishMessage(messageBody []byte, queueName string) {
	createQueue(publisherChannel, queueName)

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
		log.Println("[RabbitMQ] Cannot publish message to RabbitMQ queue. error: ", err)
	}
}

func GetConsumer(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	createQueue(channel, queueName)

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
		log.Println("[RabbitMQ] Cannot get RabbitMQ consumer channel. error: ", err)
	}
	return messagesConsumer, err
}
