package kafka

import (
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
	topic := "test006"
	err = c.Consume(topic, subscribeAvroFunc)
	if err != nil {
		panic(err.Error())
	}

}

func subscribeAvroFunc(key interface{}, value interface{}) {

	var msgValue record.ExampleRequest
	//注意传参为record的地址
	err := DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("msg => key=%s, value=%v\n", string(key.([]byte)), msgValue)
}
