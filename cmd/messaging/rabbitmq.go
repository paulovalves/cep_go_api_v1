package messaging

import (
	"context"
	"fmt"
	"log"
	"service"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	exchangeName = "events-exchange"
	queueName    = "events-queue"
)

func Connect(ctx context.Context) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %v", err)
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind a queue: %v", err)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	for d := range msgs {
		if d.Body == nil {
			log.Printf("Received an empty message")
			continue
		}
		log.Printf("Received a message: %v", d.Body)
		messagingService := service.ProcessMessage(d.Body)
		if messagingService.Error != nil {
			log.Printf("Error while processing message: %v", messagingService.Error)
			continue
		}
		messageBytes := []byte(messagingService.Message)
		// Publish the response to the exchange
		err := PublishMessage(ch, messageBytes)
		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
		}
	}

	return nil
}

func PublishMessage(ch *amqp.Channel, message []byte) error {
	err := ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}
	return nil
}
