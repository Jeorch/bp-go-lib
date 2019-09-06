package test

import (
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/utils"
	"testing"
)

func TestConvert_JsonFile2Map(t *testing.T) {
	filePath := "../resources/kafka_config.json"
	m, err := utils.Convert_JsonFile2Map(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
