package log

import (
	"github.com/PharbersDeveloper/bp-go-lib/test"
	"testing"
	"time"
)

func TestLogicLogger_Info(t *testing.T) {

	//设置项目范围内的环境变量
	test.SetEnv()

	jobId := "job-001"
	traceId := "trace-001"
	userId := "user-001"

	NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Info("aaa")
	NewLogicLoggerBuilder().Build().Info("aaa")
	NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Infof("aaa=%s", "aaa")
	NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Infoln("aaa", "aaa")

	NewLogicLoggerBuilder().SetJobId(jobId).SetUserId(userId).Build().Trace("ttt")

	bLogger := NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build()
	bLogger.Info("bbb")
	bLogger.Info("ccc")
}

func TestLogicLoggerRolling(t *testing.T) {
	//设置项目范围内的环境变量
	test.SetEnv()

	jobId := "job-001"
	traceId := "trace-001"
	userId := "user-001"

	tic := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tic.C:
			NewLogicLoggerBuilder().SetJobId(jobId).SetTraceId(traceId).SetUserId(userId).Build().Info("aaa")
			NewLogicLoggerBuilder().Build().Info("bbb")
		}
	}
}
