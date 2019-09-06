package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	defaultTimestampFormat = time.RFC3339
)

type logicLogger struct {
	jobId       string
	traceId     string
	userId      string
}

func initLogicLogger() {

	// SetReportCaller sets whether the standard logger will include the calling
	logrus.SetReportCaller(true)

	// Set logic logger formatter
	// 根据项目环境变量 $BP_LOG_TIME_FORMAT => 时间格式（默认time.RFC3339）
	timeFormat := os.Getenv("BP_LOG_TIME_FORMAT")
	switch timeFormat {
	case "":
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat:defaultTimestampFormat})
	default:
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat: timeFormat})
	}

	//根据项目环境变量 $BP_LOG_OUTPUT => 打印还是写文件(默认打印)
	logOutput := os.Getenv("BP_LOG_OUTPUT")
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

	//根据项目环境变量 $BP_LOG_LEVEL => 设置日志等级（默认trace）
	logLevel := os.Getenv("BP_LOG_LEVEL")
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

// NewLogicLoggerBuilder => Generate a LogicLogger.
// Require env $PROJECT_NAME $BP_LOG_TIME_FORMAT $BP_LOG_OUTPUT $BP_LOG_LEVEL
func NewLogicLoggerBuilder() (lg *logicLogger) {
	initLogicLogger()
	return new(logicLogger)
}

func (lg *logicLogger) Build() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{"JobId": lg.jobId, "TraceId": lg.traceId, "UserId": lg.userId})
}

func (lg *logicLogger) SetJobId(jobId string) *logicLogger {
	lg.jobId = jobId
	return lg
}

func (lg *logicLogger) SetTraceId(traceId string) *logicLogger {
	lg.traceId = traceId
	return lg
}

func (lg *logicLogger) SetUserId(userId string) *logicLogger {
	lg.userId = userId
	return lg
}

//func (lg *logicLogger) SetFormatter(formatter logrus.Formatter) {
//	logrus.SetFormatter(formatter)
//}

// AddHook adds a hook to the standard logger hooks.
func (lg *logicLogger) AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

