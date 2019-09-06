package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Convert_JsonFile2Map(filePath string) (m map[string]interface{}, err error) {
	b, _ := ioutil.ReadFile(filePath)
	if err = json.Unmarshal(b, &m); err != nil {
		err = fmt.Errorf("bp-go-lib.utils.Convert_JsonFile2Map error! filePath=%s, error msg=%s", filePath, err.Error())
		return
	}
	return
}
