package kafka

import (
	"bytes"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record"
	"github.com/PharbersDeveloper/bp-go-lib/test"
	"testing"
)

func TestConsume(t *testing.T) {

	test.SetEnv()

	//c, err := NewKafkaBuilder().BuildConsumer()
	c, err := NewKafkaBuilder().SetGroupId("test-group-001").BuildConsumer()
	if err != nil {
		panic(err.Error())
	}
	err = c.Consume("test", subscribeFunc)
	if err != nil {
		panic(err.Error())
	}

}

func subscribeFunc(key interface{}, value interface{}) {
	fmt.Printf("subscribeFunc => key=%s, value=%s\n", string(key.([]byte)), string(value.([]byte)))
}

func TestConsumeAvro(t *testing.T) {

	test.SetEnv()

	//c, err := NewKafkaBuilder().BuildConsumer()
	c, err := NewKafkaBuilder().SetGroupId("test-group-001").BuildConsumer()
	if err != nil {
		panic(err.Error())
	}
	err = c.Consume("test001", subscribeAvroFunc)
	if err != nil {
		panic(err.Error())
	}

}

func subscribeAvroFunc(key interface{}, value interface{}) {

	buf := bytes.NewBuffer(value.([]byte))

	newDemoStruct, err := record.DeserializeExampleRequest(buf)
	if err != nil {
		fmt.Printf("Error deserializing struct: %v\n", err)
		return
	}

	fmt.Printf("subscribeFunc => key=%s, value=%v\n", string(key.([]byte)), newDemoStruct)
}
