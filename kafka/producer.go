package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type BpProducer struct {
	producer *kafka.Producer
}

func (bpb *BpKafkaBuilder) BuildProducer() (*BpProducer, error) {
	p, err := kafka.NewProducer(bpb.config)
	if err != nil {
		return nil, err
	}
	bpp := new(BpProducer)
	bpp.producer = p
	return bpp, err
}

func (bpp *BpProducer) Produce(topic string, key []byte, value []byte) (err error) {

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

	// 还是要关闭producer不然会占用资源（200-500+顺序压测出现线程无法调用情况），我曾在调用方写单例发现库中封装太好了，啥都取不到
	bpp.producer.Close()
	close(deliveryChan)

	return
}
