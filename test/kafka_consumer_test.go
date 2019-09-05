package test

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestConsume(t *testing.T) {

	err := consume("test", subscribeFunc)
	if err != nil {
		panic(err.Error())
	}

}

func subscribeFunc(key interface{}, value interface{}) {
	fmt.Printf("subscribeFunc => key=%s, value=%s\n", string(key.([]byte)), string(value.([]byte)))
	//time.Sleep(10 * time.Second)
	fmt.Println("subscribeFunc DONE!")
}

func consume(topic string, subscribeFunc func(interface{}, interface{}) ) (err error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{

		"bootstrap.servers": "123.56.179.133:9092",
		"security.protocol": "SSL",	//默认使用SSL
		"ssl.ca.location": "/Users/jeorch/kit/kafka-secrets/snakeoil-ca-1.crt",
		"ssl.certificate.location": "/Users/jeorch/kit/kafka-secrets/kafkacat-ca1-signed.pem",
		"ssl.key.location": "/Users/jeorch/kit/kafka-secrets/kafkacat.client.key",
		"ssl.key.password": "pharbers",
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              "test20190828",
		"session.timeout.ms":    6000,
		//"auto.offset.reset":        "earliest",
		"auto.offset.reset":        "latest",
	})
	if err != nil {
		return
	}

	err = c.Subscribe(topic, nil)
	if err != nil {
		return
	}

	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				subscribeFunc(e.Key, e.Value)
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				continue
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

	return
}
