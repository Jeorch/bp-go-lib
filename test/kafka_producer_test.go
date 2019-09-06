package test

import (
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"testing"
)

func TestProduce(t *testing.T) {

	setEnv()

	p, err := kafka.NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}
	err = p.Produce("test", []byte("key001"), []byte("value001"))
	if err != nil {
		panic(err.Error())
	}

}


