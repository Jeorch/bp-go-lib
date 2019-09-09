// Package kafka is bp-go-lib's kafka middleware.
package kafka

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/PharbersDeveloper/bp-go-lib/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

const (
	groupIdKey = "group.id"
)

type BpKafkaBuilder struct {
	config *kafka.ConfigMap
}

func NewKafkaBuilder() *BpKafkaBuilder {
	c, err := generateKafkaConfig()
	if err != nil {
		panic(err.Error())
	}
	bpb := new(BpKafkaBuilder)
	bpb.config = c
	return bpb
}

// SetConfig set kafka config by ConfigMap.
// ConfigMap is a map contaning standard librdkafka configuration properties as documented in:
// https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
func (bpb *BpKafkaBuilder) SetConfig(configMap *kafka.ConfigMap) *BpKafkaBuilder {
	bpb.config = configMap
	return bpb
}

// AddConfig add kafka config by ConfigMap.
// ConfigMap is a map contaning standard librdkafka configuration properties as documented in:
// https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
func (bpb *BpKafkaBuilder) AddConfig(configMap *kafka.ConfigMap) *BpKafkaBuilder {
	for k, v := range *configMap {
		err := bpb.config.SetKey(k, v)
		if err != nil {
			panic(err.Error())
		}
	}
	return bpb
}

func (bpb *BpKafkaBuilder) SetGroupId(id string) *BpKafkaBuilder {
	err := bpb.config.SetKey(groupIdKey, id)
	if err != nil {
		panic(err.Error())
	}
	return bpb
}

func generateKafkaConfig() (*kafka.ConfigMap, error) {
	//根据项目环境变量设置 kafka config map
	filePath := os.Getenv(env.KafkaConfigPath)
	if filePath == "" {
		return nil, fmt.Errorf("no kafka config file path set in %s env", env.KafkaConfigPath)
	}

	m, err := utils.Convert_JsonFile2Map(filePath)
	if err != nil {
		return nil, err
	}

	// ConfigMap is a map contaning standard librdkafka configuration properties as documented in:
	// https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
	config := &kafka.ConfigMap{}
	for k, v := range m {
		err = config.SetKey(k, v)
		if err != nil {
			return nil, err
		}
	}
	return config, err
}
