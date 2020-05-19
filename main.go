package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func send(ch *amqp.Channel, q *amqp.Queue, msg string) {
	body := "Olá mundo"
	err := ch.Publish(
		"",     //exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Falha ao publicar mensagem")
	log.Printf("Enviado %s", body)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err, "Falha ao conectar")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Falha ao abrir canal")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", //name
		false,   //durable
		false,   //delete when unused
		false,   //exclusive
		false,   //no-wait
		nil,     // arguments
	)
	failOnError(err, "Falha ao declarar fila")

	body := "Olá mundo"
	send(ch, &q, body)
}
