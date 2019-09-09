package utils

import (
	"fmt"
	"testing"
)

func TestConvert_JsonFile2Map(t *testing.T) {
	filePath := "../resources/kafka_config.json"
	m, err := Convert_JsonFile2Map(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
