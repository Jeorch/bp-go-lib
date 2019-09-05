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
	userId := "trace-001"

	log.NewLogicLogger().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Info("aaa")
	log.NewLogicLogger().Info("aaa")
	log.NewLogicLogger().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Infof("aaa=%s", "aaa")
	log.NewLogicLogger().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Infoln("aaa", "aaa")

	log.NewLogicLogger().SetJobId(jobId).SetUserId(userId).Trace("ttt")

	bLogger := log.NewLogicLogger()
	bLogger.SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Info("bbb")
	bLogger.Info("bbb")
}

func setEnv() {
	//项目范围内的环境变量
	_ = os.Setenv("PROJECT_NAME", "bp-go-lib")
	_ = os.Setenv("BP_LOG_TIME_FORMAT", "2006-01-02 15:04:05")
	_ = os.Setenv("BP_LOG_OUTPUT", "console")
	_ = os.Setenv("BP_LOG_LEVEL", "info")
}