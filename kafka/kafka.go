package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

const (
	GroupId = "group.id"
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

func (bpb *BpKafkaBuilder) SetConfig(configMap *kafka.ConfigMap, useConfigFile bool) *BpKafkaBuilder {
	if !useConfigFile {
		bpb.config = configMap
	} else {
		for k, v := range *configMap {
			err := bpb.config.SetKey(k, v)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	return bpb
}

func (bpb *BpKafkaBuilder) SetGroupId(groupId string) *BpKafkaBuilder {
	err := bpb.config.SetKey(GroupId, groupId)
	if err != nil {
		panic(err.Error())
	}
	return bpb
}
