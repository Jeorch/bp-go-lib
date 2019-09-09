package log

import (
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	defaultTimestampFormat = time.RFC3339
)

type logicLoggerBuilder struct {
	jobId       string
	traceId     string
	userId      string
}

func initLogicLogger() {

	// SetReportCaller sets whether the standard logger will include the calling
	logrus.SetReportCaller(true)

	// Set logic logger formatter
	// 根据项目环境变量设置时间格式（默认time.RFC3339）
	timeFormat := os.Getenv(env.LogTimeFormat)
	switch timeFormat {
	case "":
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat:defaultTimestampFormat})
	default:
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat: timeFormat})
	}

	//根据项目环境变量设置打印还是写文件(默认打印)
	logOutput := os.Getenv(env.LogOutput)
	switch logOutput {
	case "", "console":
		logrus.SetOutput(os.Stdout)
	default:
		file, err := os.OpenFile(logOutput, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err.Error())
		}
		logrus.SetOutput(file)
	}

	//根据项目环境变量设置日志等级（默认trace）
	logLevel := os.Getenv(env.LogLevel)
	switch logLevel {
	case "":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		lvl, err := logrus.ParseLevel(logLevel)
		if err != nil {
			panic(err.Error())
		}
		logrus.SetLevel(lvl)
	}

}

// NewLogicLoggerBuilder generate a logicLoggerBuilder.
func NewLogicLoggerBuilder() (lg *logicLoggerBuilder) {
	initLogicLogger()
	return new(logicLoggerBuilder)
}

func (lg *logicLoggerBuilder) Build() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{"JobId": lg.jobId, "TraceId": lg.traceId, "UserId": lg.userId})
}

func (lg *logicLoggerBuilder) SetJobId(jobId string) *logicLoggerBuilder {
	lg.jobId = jobId
	return lg
}

func (lg *logicLoggerBuilder) SetTraceId(traceId string) *logicLoggerBuilder {
	lg.traceId = traceId
	return lg
}

func (lg *logicLoggerBuilder) SetUserId(userId string) *logicLoggerBuilder {
	lg.userId = userId
	return lg
}

//func (lg *logicLoggerBuilder) SetFormatter(formatter logrus.Formatter) {
//	logrus.SetFormatter(formatter)
//}

// AddHook adds a hook to the standard logger hooks.
func (lg *logicLoggerBuilder) AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

