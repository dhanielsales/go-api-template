package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dhanielsales/go-api-template/internal/config/env"

	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/dhanielsales/go-api-template/pkg/messaging"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (kaf *KafkaProvider) errorHandler(err error, message *kafka.Message) {
	topic := *message.TopicPartition.Topic
	partition := message.TopicPartition.Partition
	offset := message.TopicPartition.Offset

	if err == nil {
		logger.Error("unknown error occurred: error is nil",
			logger.LogField("topic", topic),
			logger.LogField("partition", partition),
			logger.LogField("offset", offset),
		)

		return
	}

	headers, err := updateErrorHeader(message, err)
	if err != nil {
		logger.Error(errors.Join(errors.New("Error upserting error payload"), err).Error(),
			logger.LogField("topic", topic),
			logger.LogField("partition", partition),
			logger.LogField("offset", offset),
		)

		return
	}

	retryTopic := getRetryTopicEvent().String()

	kaf.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &retryTopic, Partition: kafka.PartitionAny},
		Value:          message.Value,
		Headers:        headers,
	}, nil)
}

func (kaf *KafkaProvider) resolveHandler(message *kafka.Message) {
	topic := *message.TopicPartition.Topic
	partition := message.TopicPartition.Partition
	offset := message.TopicPartition.Offset

	headers, err := resolveErrorHeader(message)
	if err != nil {
		logger.Error(errors.Join(errors.New("Error resolving error payload"), err).Error(),
			logger.LogField("topic", topic),
			logger.LogField("partition", partition),
			logger.LogField("offset", offset),
		)

		return
	}

	retryTopic := getRetryTopicEvent().String()
	kaf.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &retryTopic, Partition: kafka.PartitionAny},
		Value:          message.Value,
		Headers:        headers,
	}, nil)
}

func resolveErrorHeader(message *kafka.Message) ([]kafka.Header, error) {
	headers := message.Headers
	result := make([]kafka.Header, 0, len(headers))

	for _, header := range headers {
		if header.Key == messaging.ERROR_HEADER_KEY {
			errorPayload := messaging.ErrorPayload{}
			err := json.Unmarshal(header.Value, &errorPayload)
			if err != nil {
				return nil, err
			}

			now := time.Now()
			errorPayload.IsResolved = true
			errorPayload.ResolvedAt = &now

			errorPayloadBytes, err := json.Marshal(errorPayload)
			if err != nil {
				return nil, err
			}

			result = append(result, kafka.Header{Key: messaging.ERROR_HEADER_KEY, Value: errorPayloadBytes})
		} else {
			result = append(result, header)
		}
	}

	return result, nil
}

func updateErrorHeader(message *kafka.Message, err error) ([]kafka.Header, error) {
	headers := message.Headers
	result := make([]kafka.Header, 0, len(headers))

	for _, header := range headers {
		if header.Key == messaging.ERROR_HEADER_KEY {
			return headers, nil
		}

		result = append(result, header)
	}

	errorPayload := messaging.ErrorPayload{
		Occourrences:         1,
		IsResolved:           false,
		OriginalOccurrenceAt: time.Now(),
		OriginalOffset:       int64(message.TopicPartition.Offset),
		ResolvedAt:           nil,
		PreviousTopic:        *message.TopicPartition.Topic,
		StackTrace:           fmt.Sprintf("%+v", err), // TODO colocar error completo, junto com stack trace
		Error:                err.Error(),
	}

	errorPayloadBytes, err := json.Marshal(errorPayload)
	if err != nil {
		return nil, err
	}

	result = append(result, kafka.Header{Key: messaging.ERROR_HEADER_KEY, Value: errorPayloadBytes})

	return result, nil
}

func errorBoundary(handler messaging.Handler) messaging.Handler {
	return func(payload messaging.EventPayload) (err error) {
		defer func() {
			if recovErr := recover(); recovErr != nil {
				if e, ok := recovErr.(error); ok {
					err = e
				} else {
					err = errors.New(fmt.Sprintf("Unknown error occurred while processing message: %s", recovErr))
				}
			}
		}()

		err = handler(payload)
		return err
	}
}

func getRetryTopicEvent() messaging.Event {
	envVars := env.GetInstance()

	return messaging.Event(fmt.Sprintf("%s.%s.%s", envVars.APP_NAME, envVars.KAFKA_GROUP_ID, messaging.RETRY_TOPIC))
}
