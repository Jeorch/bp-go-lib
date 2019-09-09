package kafka

import (
	"bytes"
	"fmt"
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

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}

	requestRecord := record.ExampleRequest{
		JobId:   "job-001",
		Tag:     "MAX",
		Configs: []string{
			"config-1",
			"config-2",
		},
	}

	// Serialize the record to a byte buffer
	var buf bytes.Buffer
	fmt.Printf("Serializing struct: %#v\n", requestRecord)
	err = requestRecord.Serialize(&buf)
	if err != nil {
		panic(err.Error())
	}

	err = p.Produce("test001", []byte("avro-key002"), buf.Bytes())
	if err != nil {
		panic(err.Error())
	}

}
