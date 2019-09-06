package test

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"testing"
)

func TestConsume(t *testing.T) {

	setEnv()

	c, err := kafka.NewKafkaBuilder().BuildConsumer()
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
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}
