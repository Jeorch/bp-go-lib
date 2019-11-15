package kafka

import (
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record"
	"github.com/PharbersDeveloper/bp-go-lib/test"
	"testing"
)

func TestProduce(t *testing.T) {

	test.SetEnv()

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}
	err = p.Produce("test", []byte("key002"), []byte("value002"))
	if err != nil {
		panic(err.Error())
	}

}

func TestProduceAvro(t *testing.T) {

	test.SetEnv()

	topic := "test006"
	requestRecord := record.ExampleRequest{
		JobId:   "job-001",
		Tag:     "MAX",
		Configs: []string{
			"config-101",
			"config-202",
		},
	}

	specificRecordByteArr, err := EncodeAvroRecord(&requestRecord)
	if err != nil {
		panic(err.Error())
	}

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}

	err = p.Produce(topic, []byte("avro-key003"), specificRecordByteArr)
	if err != nil {
		panic(err.Error())
	}

}
