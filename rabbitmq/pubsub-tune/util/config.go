package util

import "github.com/spf13/viper"

type Config struct {
	MQ struct {
		Url          string `mapstructure:"url"`
		ExchangeName string `mapstructure:"exchange_name"`
		ExchangeType string `mapstructure:"exchange_type"`
	} `mapstructure:"mq"`
	Subscribers struct {
		Connections           int    `mapstructure:"connections"`
		ChannelsPerConnection int    `mapstructure:"channels_per_connection"`
		QueuesPerChannel      int    `mapstructure:"queues_per_channel"`
		PrefetchCount         int    `mapstructure:"prefetch_count"`
		RoutingKey            string `mapstructure:"routing_key"`
		AutoAck               bool   `mapstructure:"auto_ack"`
		QueueNamePrefix       string `mapstructure:"queue_name_prefix"`
	} `mapstructure:"subscribers"`
	Publisher struct {
		RoutingKey        string `mapstructure:"routing_key"`
		DataSize          int    `mapstructure:"data_size"`
		PackagesPerSecond int    `mapstructure:"packages_per_second"`
	} `mapstructure:"publisher"`
}

func LoadConfig(configName string) (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return
		} else {
			// Config file was found but another error was produced
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
