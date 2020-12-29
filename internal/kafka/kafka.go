package kafka

import (
	"context"
	"errors"
	"fmt"
	"my-github/clean-code-microservice-golang/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type InternalKafkaImpl struct {
	kp *kafka.Producer
}

type InternalKafka interface {
	Publish(ctx context.Context, key, message []byte) (err error)
}

func NewInternalKafka(kp *kafka.Producer) InternalKafka {
	return &InternalKafkaImpl{kp: kp}
}

func (i *InternalKafkaImpl) Publish(ctx context.Context, key, message []byte) (err error) {
	var (
		topics       = config.GetStringSlice("topics.kafka")
		topic        string
		deliveryChan = make(chan kafka.Event)
	)

	// check if topics has a value
	if len(topics) > 0 {
		// set topic to topics index[0]
		topic = topics[0]
	}

	// Call produce message from kafka confluent
	// set kafka message such as partition+topic, value message, and key
	// set receive channel so when producer has success it returns to delivery channel
	err = i.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Key:            key,
	}, deliveryChan)

	e := <-deliveryChan

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return errors.New(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	}
	close(deliveryChan)
	return
}
