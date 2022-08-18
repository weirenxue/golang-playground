package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"rabbitmq/util"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	conn *amqp.Connection
}

func New(url string) (*rabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	r := rabbitMQ{
		conn: conn,
	}
	return &r, nil
}

func (r *rabbitMQ) newChannel() (*rabbitMQChannel, error) {
	return NewMQChannel(r.conn)
}

type rabbitMQChannel struct {
	channel *amqp.Channel
}

func NewMQChannel(conn *amqp.Connection) (*rabbitMQChannel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQChannel{
		channel: ch,
	}, nil
}

func (c *rabbitMQChannel) ExchangeDeclare(name, _type string) error {
	return c.channel.ExchangeDeclare(
		name,  // name
		_type, // type
		false, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
}

func (c *rabbitMQChannel) Publish(exchangeName string, routingKey string, data []byte) error {
	return c.channel.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Body: data,
		},
	)
}

var alphanum = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var str strings.Builder
	var k = len(alphanum)
	for i := 0; i < n; i++ {
		str.WriteByte(alphanum[rand.Intn(k)])
	}
	return str.String()
}

func main() {
	config, err := util.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}
	configJsonValue, _ := json.MarshalIndent(config, "", "    ")
	log.Println(string(configJsonValue))

	data := []byte(RandomString(config.Publisher.DataSize))

	r, err := New(config.MQ.Url)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := r.newChannel()
	if err != nil {
		log.Fatal(err)
	}
	if err := ch.ExchangeDeclare(config.MQ.ExchangeName, config.MQ.ExchangeType); err != nil {
		log.Fatal(err)
	}
	timer := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-timer.C:
			for i := 0; i < config.Publisher.PackagesPerSecond; i++ {
				go ch.Publish(config.MQ.ExchangeName, config.Publisher.RoutingKey, data)
			}
		}
	}

}
