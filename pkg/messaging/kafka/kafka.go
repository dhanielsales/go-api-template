package kafka

import (
	"github.com/dhanielsales/go-api-template/pkg/messaging"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type topicType = string
type groupType = string

type KafkaProvider struct {
	notify        chan struct{}
	producer      *kafka.Producer
	consumers     map[groupType]*kafka.Consumer
	subscriptions map[groupType]map[topicType]messaging.Handler
}

func New() (*KafkaProvider, error) {
	return &KafkaProvider{
		subscriptions: make(map[groupType]map[topicType]messaging.Handler),
		consumers:     make(map[groupType]*kafka.Consumer),
		notify:        make(chan struct{}),
	}, nil
}

func (kaf *KafkaProvider) Start() error {
	err := kaf.setupConsumers()
	if err != nil {
		return err
	}

	err = kaf.setupProducer()
	if err != nil {
		return err
	}

	go kaf.watchMesseges()

	for group := range kaf.consumers {
		go kaf.consumeMessages(group)
	}

	return nil
}

func (kaf *KafkaProvider) Cleanup() {
	kaf.notify <- struct{}{}
}

func (kaf *KafkaProvider) Subscribe(event messaging.Event, handler messaging.Handler) error {
	topic := event.String()
	group := event.Group()

	if kaf.subscriptions[group] == nil {
		kaf.subscriptions[group] = make(map[topicType]messaging.Handler)
	}

	if kaf.subscriptions[group][topic] != nil {
		return messaging.ErrEventAlreadySubscribed
	}

	kaf.subscriptions[group][topic] = handler

	return nil
}

func (kaf *KafkaProvider) SubscribeFanOut(event messaging.Event, handler messaging.Handler) error {
	topic := event.String()
	group := messaging.FANOUT_GROUP

	if kaf.subscriptions[group] == nil {
		kaf.subscriptions[group] = make(map[topicType]messaging.Handler)
	}

	if kaf.subscriptions[group][topic] != nil {
		return messaging.ErrEventAlreadySubscribed
	}

	kaf.subscriptions[group][topic] = handler

	return nil
}

func (kaf *KafkaProvider) Publish(event messaging.Event, data []byte, meta messaging.Meta, errPayload *messaging.ErrorPayload) error {
	topic := event.String()

	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}

	if meta != nil {
		metaJSON, err := meta.ToJSON()
		if err != nil {
			return err
		}

		message.Headers = []kafka.Header{{Key: messaging.META_HEADER_KEY, Value: metaJSON}}
	}

	if errPayload != nil {
		errPayloadJSON, err := errPayload.ToJSON()
		if err != nil {
			return err
		}

		message.Headers = append(message.Headers, kafka.Header{Key: messaging.ERROR_HEADER_KEY, Value: errPayloadJSON})
	}

	err := kaf.producer.Produce(&message, nil)
	if err != nil {
		return err
	}

	return nil
}
