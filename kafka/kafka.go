// Package kafka is bp-go-lib's kafka middleware.
package kafka

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/PharbersDeveloper/bp-go-lib/utils"
	"github.com/actgardner/gogen-avro/container"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	goAvro "github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"os"
	"sync"
)

const (
	groupIdKey = "group.id"
)

var (
	once sync.Once
	conf *kafka.ConfigMap
)

type BpKafkaBuilder struct {
	config *kafka.ConfigMap
}

func NewKafkaBuilder() *BpKafkaBuilder {
	once.Do(generateKafkaConfig)
	bpb := new(BpKafkaBuilder)
	bpb.config = conf
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

func generateKafkaConfig() {
	//根据项目环境变量设置 kafka config map
	filePath := os.Getenv(env.KafkaConfigPath)
	if filePath == "" {
		panic(fmt.Sprintf("no kafka config file path set in %s env", env.KafkaConfigPath))
	}

	m, err := utils.Convert_JsonFile2Map(filePath)
	if err != nil {
		panic(err.Error())
	}

	// ConfigMap is a map contaning standard librdkafka configuration properties as documented in:
	// https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
	conf = &kafka.ConfigMap{}
	for k, v := range m {
		err = conf.SetKey(k, v)
		if err != nil {
			panic(err.Error())
		}
	}
}

func EncodeAvroRecord(obj container.AvroRecord) ([]byte, error) {

	var magic_bytes = []byte{0}
	schemaRegistryUrl := os.Getenv(env.KafkaSchemaRegistryUrl)
	if schemaRegistryUrl == "" {
		panic(fmt.Sprintf("no kafka config file path set in %s env", env.KafkaSchemaRegistryUrl))
	}

	if obj == nil {
		return nil, nil
	}

	schema, err := goAvro.ParseSchema(obj.Schema())
	if err != nil {
		panic(err.Error())
	}

	subject := schema.GetName() + "-value"
	//注册schema
	id, err := kafkaAvro.NewCachedSchemaRegistryClientAuth(schemaRegistryUrl, nil).Register(subject, schema)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	_, err = buffer.Write(magic_bytes)
	if err != nil {
		return nil, err
	}
	idSlice := make([]byte, 4)
	binary.BigEndian.PutUint32(idSlice, uint32(id))
	_, err = buffer.Write(idSlice)
	if err != nil {
		return nil, err
	}

	enc := goAvro.NewBinaryEncoder(buffer)
	var writer goAvro.DatumWriter
	writer = goAvro.NewSpecificDatumWriter()
	writer.SetSchema(schema)
	err = writer.Write(obj, enc)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DecodeAvroRecord(bytes []byte, obj interface{}) error {
	schemaRegistryUrl := os.Getenv(env.KafkaSchemaRegistryUrl)
	if schemaRegistryUrl == "" {
		panic(fmt.Sprintf("no kafka config file path set in %s env", env.KafkaSchemaRegistryUrl))
	}

	decoder := kafkaAvro.NewKafkaAvroDecoder(schemaRegistryUrl)

	//此处的obj应为record的指针
	err := decoder.DecodeSpecific(bytes, obj)

	return err
}
