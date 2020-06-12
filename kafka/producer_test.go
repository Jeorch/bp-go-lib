package kafka

import (
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka/record/PhEventMsg"
	"github.com/PharbersDeveloper/bp-go-lib/test"
	"testing"
)

func TestProduce(t *testing.T) {

	test.SetEnv()

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}
	err = p.Produce("test", []byte("key003"), []byte("value003"))
	if err != nil {
		panic(err.Error())
	}

}

func TestProduceAvro(t *testing.T) {

	test.SetEnv()

	topic := "oss_msg_gen_cube"

	json, err := json.Marshal(map[string]string{
		"InputDataType": "hive",
		"InputPath": "SELECT * FROM result",
		"OutputDataType": "es",
		"OutputPath": "ekscube2",
		"strategy": "handle-hive-result",
	})
	if err != nil {
		panic(err.Error())
	}

	requestRecord := PhEventMsg.EventMsg{
		JobId:   "testId",
		TraceId: "testId",
		Type:    "GenCube-Start",
		Data:    string(json),
	}

	specificRecordByteArr, err := EncodeAvroRecord(&requestRecord)
	if err != nil {
		panic(err.Error())
	}

	p, err := NewKafkaBuilder().BuildProducer()
	if err != nil {
		panic(err.Error())
	}

	err = p.Produce(topic, []byte("testId"), specificRecordByteArr)
	if err != nil {
		panic(err.Error())
	}

}
