package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type bpProducer struct {
	producer *kafka.Producer
}

func (bpb *BpKafkaBuilder) BuildProducer() (*bpProducer, error) {
	p, err := kafka.NewProducer(bpb.config)
	if err != nil {
		return nil, err
	}
	bpp := new(bpProducer)
	bpp.producer = p
	return bpp, err
}

func (bpp *bpProducer) Produce(topic string, key []byte, value []byte) (err error) {

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:          	key,
		Value:          value,
	}

	err = bpp.producer.Produce(&msg, deliveryChan)
	if err != nil {
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		err = m.TopicPartition.Error
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return
}
