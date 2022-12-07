# Publish/Subscribe Tune

The goal of this project is to know when RabbitMQ traffic is blocking in topic mode.

## How can we use this project?

We can adjust some parameters in `config.toml` and watch the traffic on the RabbitMQ [management page](http://localhost:15672/).

There are two main command  

1. `cmd/subscribers/main.go`: This command is simulating one or many subscribers. We can modify configuration of `subscribers` block in `config.toml` to adjust the action of this command. The configuration instructions are as follows

    1. `connections`: Total number of connections.
    1. `channels_per_connection`: Number of channels per connection. The total number of channels is `connections * channels_per_connection`
    1. `queues_per_channel`: Number of queues per channel. The total number of queues is `connections * channels_per_connection * queues_per_channel`
    1. `prefetch_count`: How many prefetches per channel.
    1. `routing_key`: What routing key to bind for every queue.
    1. `auto_ack`: Whether to use auto ack or not. If true, prefetching will not work.
    1. `queue_name_prefix`: Each queue name is `${queue_name_prefix}-${i-th connection}-${j-th channel}-${k-th queue}`. E.g.: `for_test-1-1-1`

1. `cmd/publisher/main.go`: This command is simulating one publisher. We can modify configuration of `publisher` block in `config.toml` to adjust the action of this command. The configuration instructions are as follows

    1. `routing_key`: What routing key to publish by
    1. `data_size`: How big a payload is, in bytes.
    1. `packages_per_second`: How many payloads are expected to be sent per second.

If your need to modify RabbitMQ connection configuration, the `mq` block in `config.toml` is for you.

## How to run?

There is no order to run `subscribers` and `publisher`.

Run the `subscribers`

```sh
bash sub.sh
```

Run the `publisher`

```sh
bash pub.sh
```
