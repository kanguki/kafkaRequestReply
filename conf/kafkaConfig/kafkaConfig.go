package kafConfig

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var DefaultTopicToProduce string = "test"
var DefaultTopicsToConsume []string = []string{"test"}

var KafkaProducerConfig = &kafka.ConfigMap{
	"bootstrap.servers":     "localhost:9092",
	"request.required.acks": -1,
	"message.max.bytes":     1000000,
}

var KafkaConsumerConfig = &kafka.ConfigMap{
	"bootstrap.servers":        "localhost:9092",
	"client.id":                "xxx",
	"group.id":                 "xxx",
	"auto.offset.reset":        "latest",
	"enable.auto.offset.store": true,
}
