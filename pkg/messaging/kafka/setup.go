package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/dhanielsales/go-api-template/internal/config/env"
	"github.com/dhanielsales/go-api-template/pkg/logger"
)

func (kaf *KafkaProvider) setupConsumers() error {
	for group, subscriptionsPerGroup := range kaf.subscriptions {
		topics := MapKeys(subscriptionsPerGroup)

		consumer, err := newConsumer(group)
		if err != nil {
			logger.Error("Consumer creation failed", logger.LogField("group", group))

			return err
		}

		kaf.consumers[group] = consumer

		err = consumer.SubscribeTopics(topics, nil)
		if err != nil {
			logger.Error("Topic subscription failed", logger.LogField("group", group))

			return err
		}
	}

	return nil
}

func (kaf *KafkaProvider) setupProducer() error {
	producer, err := newProducer()
	if err != nil {
		return err
	}

	kaf.producer = producer

	return nil
}

func newProducer() (*kafka.Producer, error) {
	envVars := env.GetInstance()

	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": envVars.KAFKA_BROKER,
		"client.id":         envVars.KAFKA_CLIENT_ID,
		"log_level":         5,
	}

	if envVars.KAFKA_SSL {
		producerConfig.SetKey("security.protocol", "SASL_SSL")
		producerConfig.SetKey("sasl.mechanism", "SCRAM-SHA-256")
		producerConfig.SetKey("sasl.username", envVars.KAFKA_USERNAME)
		producerConfig.SetKey("sasl.password", envVars.KAFKA_PASSWORD)
		producerConfig.SetKey("log_level", 1)
	}

	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func newConsumer(groupId string) (*kafka.Consumer, error) {
	envVars := env.GetInstance()

	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": envVars.KAFKA_BROKER,
		"client.id":         envVars.KAFKA_CLIENT_ID,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	}

	if envVars.KAFKA_SSL {
		consumerConfig.SetKey("security.protocol", "SASL_SSL")
		consumerConfig.SetKey("sasl.mechanism", "SCRAM-SHA-256")
		consumerConfig.SetKey("sasl.username", envVars.KAFKA_USERNAME)
		consumerConfig.SetKey("sasl.password", envVars.KAFKA_PASSWORD)
	}

	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
