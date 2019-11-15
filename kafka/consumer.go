package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

type bpConsumer struct {
	consumer *kafka.Consumer
}

func (bpb *bpKafkaBuilder) BuildConsumer() (*bpConsumer, error) {
	c, err := kafka.NewConsumer(bpb.config)
	if err != nil {
		return nil, err
	}
	bpp := new(bpConsumer)
	bpp.consumer = c
	return bpp, err
}

func (bpc *bpConsumer) Consume(topic string, subscribeFunc func(interface{}, interface{}) ) (err error) {

	err = bpc.consumer.Subscribe(topic, nil)
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
			ev := bpc.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				subscribeFunc(e.Key, e.Value)
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
	bpc.consumer.Close()

	return
}
