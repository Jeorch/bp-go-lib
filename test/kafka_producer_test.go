package test

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

func TestProduce(t *testing.T) {

	err := produce("test", []byte("aaa"), []byte("bbb"))
	if err != nil {
		panic(err.Error())
	}

}

func produce(topic string, key []byte, value []byte) (err error) {

	//_ = os.Setenv("BM_KAFKA_BROKER", "123.56.179.133:9092")
	//_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", "http://123.56.179.133:8081")
	//_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", "test20190828")
	//_ = os.Setenv("BM_KAFKA_CA_LOCATION", "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt")
	//_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem")
	//_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key")
	//_ = os.Setenv("BM_KAFKA_SSL_PASS", "pharbers")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "123.56.179.133:9092",
		"security.protocol": "SSL",	//默认使用SSL
		"ssl.ca.location": "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt",
		"ssl.certificate.location": "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem",
		"ssl.key.location": "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key",
		"ssl.key.password": "pharbers",
	})
	if err != nil {
		panic(err.Error())
	}

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:          	key,
		Value:          value,
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}

	err = p.Produce(&msg, deliveryChan)
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
