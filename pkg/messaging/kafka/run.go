package kafka

import (
	"errors"
	"time"

	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/dhanielsales/go-api-template/pkg/messaging"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func MapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (kaf *KafkaProvider) watchMesseges() {
	for e := range kaf.producer.Events() {
		select {
		case <-kaf.notify:
			kaf.producer.Close()
			return
		default:
			if message, ok := e.(*kafka.Message); ok {
				if message.TopicPartition.Error != nil {
					logger.Error(errors.Join(errors.New("Event delivery failed"), message.TopicPartition.Error).Error(),
						logger.LogField("topic", *message.TopicPartition.Topic),
						logger.LogField("partition", message.TopicPartition.Partition),
						logger.LogField("offset", message.TopicPartition.Offset),
					)
				}
			}
		}
	}
}

// TODO Colocar o errorHandler nos erros do consumeMessages

func (kaf *KafkaProvider) consumeMessages(group string) {
	consumer := kaf.consumers[group]
	subs := kaf.subscriptions[group]

	for {
		select {
		case <-kaf.notify:
			consumer.Close()
			return
		default:
			message, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				if kafkaError, ok := err.(kafka.Error); ok {
					if !kafkaError.IsTimeout() {
						logger.Error(errors.Join(errors.New("Event consuming failed"), kafkaError).Error())
						continue
					}
				} else {
					logger.Error(errors.Join(errors.New("Unknown error occurred while consuming event"), err).Error())
					continue
				}

				continue
			}

			topic := *message.TopicPartition.Topic
			partition := message.TopicPartition.Partition
			offset := message.TopicPartition.Offset
			headers := message.Headers

			handler := subs[topic]
			if handler == nil {
				logger.Error("Event not subscribed",
					logger.LogField("topic", topic),
					logger.LogField("partition", partition),
					logger.LogField("offset", offset),
				)
				continue
			}

			event := messaging.Event(topic)
			if err = event.Validate(); err != nil {
				logger.Error(errors.Join(errors.New("Invalid event"), err).Error(),
					logger.LogField("topic", topic),
					logger.LogField("partition", partition),
					logger.LogField("offset", offset),
				)
				continue
			}

			messagePayload := messaging.EventPayload{
				Data:  message.Value,
				Event: event,
			}

			// Get Meta
			meta := messaging.Meta{}
			for _, header := range headers {
				if header.Key == messaging.META_HEADER_KEY {
					err := meta.FromJSON(header.Value)
					if err != nil {
						logger.Error("Invalid meta header",
							logger.LogField("topic", topic),
							logger.LogField("partition", partition),
							logger.LogField("offset", offset),
						)
						continue
					}
					messagePayload.Meta = meta
					break
				}
			}

			// Get ErrorPayload
			var errorPayload messaging.ErrorPayload
			errorExistsPreviously := false
			for _, header := range headers {
				if header.Key == messaging.ERROR_HEADER_KEY {
					errorExistsPreviously = true
					err := errorPayload.FromJSON(header.Value)
					if err != nil {
						logger.Error("Invalid error payload header",
							logger.LogField("topic", topic),
							logger.LogField("partition", partition),
							logger.LogField("offset", offset),
						)
						continue
					}
					messagePayload.ErrPayload = errorPayload
					break
				}
			}

			withBoundary := errorBoundary(handler)
			err = withBoundary(messagePayload)
			if err != nil {
				logger.Error("Event handler failed",
					logger.LogField("topic", topic),
					logger.LogField("partition", partition),
					logger.LogField("offset", offset),
				)

				kaf.errorHandler(err, message)
				continue
			}

			if errorExistsPreviously && topic != getRetryTopicEvent().String() {
				kaf.resolveHandler(message)
			}

			consumer.CommitMessage(message)
		}
	}
}
