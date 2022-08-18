package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbitmq/util"

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

func (c *rabbitMQChannel) QoS(prefetchCount int) error {
	return c.channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)
}

func (c *rabbitMQChannel) QueueDeclare(name string) (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		name,  // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (c *rabbitMQChannel) QueueBind(queueName, routingKey, exchangeName string) error {
	return c.channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
	)
}

func (c *rabbitMQChannel) Consume(queueName string, autoAck bool) (<-chan amqp.Delivery, error) {
	return c.channel.Consume(
		queueName, // queue
		"",        // consumer
		autoAck,   // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}

func main() {
	config, err := util.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}
	configJsonValue, _ := json.MarshalIndent(config, "", "    ")
	log.Println(string(configJsonValue))

	for i := 1; i < config.Subscribers.Connections+1; i++ {
		r, err := New(config.MQ.Url)
		if err != nil {
			log.Fatal(err)
		}
		for j := 1; j < config.Subscribers.ChannelsPerConnection+1; j++ {
			ch, err := r.newChannel()
			if err != nil {
				log.Fatal(err)
			}
			if err := ch.ExchangeDeclare(config.MQ.ExchangeName, config.MQ.ExchangeType); err != nil {
				log.Fatal(err)
			}
			if err := ch.QoS(config.Subscribers.PrefetchCount); err != nil {
				log.Fatal(err)
			}
			for k := 1; k < config.Subscribers.QueuesPerChannel+1; k++ {
				q, err := ch.QueueDeclare(fmt.Sprintf("%s-%d-%d-%d", config.Subscribers.QueueNamePrefix, i, j, k))
				if err != nil {
					log.Fatal(err)
				}
				if err := ch.QueueBind(q.Name, config.Subscribers.RoutingKey, config.MQ.ExchangeName); err != nil {
					log.Fatal(err)
				}
				msgs, err := ch.Consume(q.Name, config.Subscribers.AutoAck)
				if err != nil {
					log.Fatal(err)
				}

				go func(i, j, k int) {
					for d := range msgs {
						if !config.Subscribers.AutoAck {
							d.Ack(false)
						}
					}
				}(i, j, k)
			}
		}
	}
	log.Println("init finish")
	select {}
}
