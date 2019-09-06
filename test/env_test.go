package test

import (
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"os"
)

func setEnv() {
	//项目范围内的环境变量
	_ = os.Setenv(env.ProjectName, "bp-go-lib")

	//log
	_ = os.Setenv(env.LogTimeFormat, "2006-01-02 15:04:05")
	_ = os.Setenv(env.LogOutput, "console")
	//_ = os.Setenv(env.LogOutput, "./bp-go-lib.log")
	_ = os.Setenv(env.LogLevel, "info")

	//kafka
	_ = os.Setenv(env.KafkaConfigPath, "../resources/kafka_config.json")

}
