package kafka

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record"
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record/PhEventMsg"
	"github.com/PharbersDeveloper/bp-go-lib/test"
	"testing"
	"time"
)

func TestConsume(t *testing.T) {

	test.SetEnv()

	//c, err := NewKafkaBuilder().BuildConsumer()
	c, err := NewKafkaBuilder().SetGroupId("test-group-001").BuildConsumer()
	if err != nil {
		panic(err.Error())
	}
	err = c.Consume("test001", subscribeFunc)
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
	c, err := NewKafkaBuilder().SetGroupId("test-group-002").BuildConsumer()
	if err != nil {
		panic(err.Error())
	}
	topic := "oss_msg_gen_cube"
	err = c.Consume(topic, eventMsgAvroFunc)
	//err = c.Consume(topic, subscribeAvroAndReplyFunc)
	if err != nil {
		panic(err.Error())
	}

}

func hiveAvroFunc(key interface{}, value interface{}) {

	var msgValue record.MapTest
	//注意传参为record的地址
	err := DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("msg => key=%s, value=%v\n", string(key.([]byte)), msgValue)
}

func eventMsgAvroFunc(key interface{}, value interface{}) {

	var msgValue PhEventMsg.EventMsg
	//注意传参为record的地址
	err := DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("msg => key=%s, value=%v\n", string(key.([]byte)), msgValue)
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

func subscribeAvroAndReplyFunc(key interface{}, value interface{}) {

	var request record.ExampleRequest
	//注意传参为record的地址
	err := DecodeAvroRecord(value.([]byte), &request)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("msg => key=%s, value=%v\n", string(key.([]byte)), request)

	topic := "test007"
	response := record.ExampleResponse{
		JobId:   request.JobId,
		Progress:     100,
		Error: "",
	}

	specificRecordByteArr, err := EncodeAvroRecord(&response)
	if err != nil {
		panic(err.Error())
	}

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}

	//假装处理了一段时间
	time.Sleep(5 * time.Second)

	err = p.Produce(topic, []byte(response.JobId), specificRecordByteArr)
	if err != nil {
		panic(err.Error())
	}

}
