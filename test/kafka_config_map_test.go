package test

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

func TestConvert_Map2KafkaConfigMap(t *testing.T) {

	filePath := "../resources/kafka_config.json"
	m, err := utils.Convert_JsonFile2Map(filePath)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(m)

	config := &kafka.ConfigMap{}
	for k, v := range m {
		err = config.SetKey(k, v)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(config)
}
