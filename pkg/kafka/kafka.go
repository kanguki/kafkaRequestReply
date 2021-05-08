package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	//import based on folder location. refer in code based on name of package.
	//here, although folder named kafkaConfig, package named kafConfig, so it's referred kafConfig
	"mo.io/kafkaReqRep/conf/kafkaConfig"
)

var KafkaProducer *kafka.Producer
var KafkaConsumer *kafka.Consumer

func InitKafka() {
	InitProducer()
	InitConsumer()
}

func Produce(message string) {
	ProduceToTopic(kafConfig.DefaultTopicToProduce, message)
}

func ProduceToTopic(topic string, message string) {
	deliveryChan := make(chan kafka.Event)
	go func() {
		for e := range deliveryChan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	pTopic := topic
	KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &pTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
}

func InitProducer() {
	kafkaProducer, err := kafka.NewProducer(kafConfig.KafkaProducerConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Producer is born.")
	KafkaProducer = kafkaProducer
}

func InitConsumer() {
	kafkaConsumer, err := kafka.NewConsumer(kafConfig.KafkaConsumerConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer is born.")
	KafkaConsumer = kafkaConsumer
	// defer KafkaConsumer.Close()
	KafkaConsumer.SubscribeTopics(kafConfig.DefaultTopicsToConsume, nil)
	go func() {
		for {
			msg, err := KafkaConsumer.ReadMessage(-1)
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()
}
