package test

import (
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"os"
	"testing"
)

func TestLogicLogger_Info(t *testing.T) {

	//设置项目范围内的环境变量
	setEnv()

	jobId := "job-001"
	traceId := "trace-001"
	userId := "user-001"

	log.NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Info("aaa")
	log.NewLogicLoggerBuilder().Build().Info("aaa")
	log.NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Infof("aaa=%s", "aaa")
	log.NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Infoln("aaa", "aaa")

	log.NewLogicLoggerBuilder().SetJobId(jobId).SetUserId(userId).Build().Trace("ttt")

	bLogger := log.NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build()
	bLogger.Info("bbb")
	bLogger.Info("ccc")
}

func setEnv() {
	//项目范围内的环境变量
	_ = os.Setenv("PROJECT_NAME", "bp-go-lib")
	_ = os.Setenv("BP_LOG_TIME_FORMAT", "2006-01-02 15:04:05")
	_ = os.Setenv("BP_LOG_OUTPUT", "console")
	//_ = os.Setenv("BP_LOG_OUTPUT", "./bp-go-lib.log")
	_ = os.Setenv("BP_LOG_LEVEL", "info")
}