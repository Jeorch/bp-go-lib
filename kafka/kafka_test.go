package kafka

import (
	"bytes"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record"
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

func TestKafkaRecord(t *testing.T) {
	// Create a new DemoSchema struct
	requestRecord := record.ExampleRequest{
		JobId:   "job-001",
		Tag:     "MAX",
		Configs: []string{
			"config-1",
			"config-2",
		},
	}

	// Serialize the struct to a byte buffer
	var buf bytes.Buffer
	fmt.Printf("Serializing struct: %#v\n", requestRecord)
	err := requestRecord.Serialize(&buf)
	if err != nil {
		panic(err.Error())
	}

	// Deserialize the byte buffer back into a struct
	newDemoStruct, err := record.DeserializeExampleRequest(&buf)
	if err != nil {
		fmt.Printf("Error deserializing struct: %v\n", err)
		return
	}
	fmt.Printf("Deserialized struct: %#v\n", newDemoStruct)
}
