package test

import (
	"github.com/PharbersDeveloper/bp-go-lib/log"
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
