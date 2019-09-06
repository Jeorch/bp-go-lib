package kafka

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/PharbersDeveloper/bp-go-lib/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

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

	config := &kafka.ConfigMap{}
	for k, v := range m {
		err = config.SetKey(k, v)
		if err != nil {
			return nil, err
		}
	}
	return config, err
}
