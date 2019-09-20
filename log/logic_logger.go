package log

import (
	"github.com/PharbersDeveloper/bp-go-lib/env"
	rollingfile "github.com/lanziliang/logrus-rollingfile-hook"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	defaultTimestampFormat   = time.RFC3339
	defaultRollingTimeFormat = "2006-01-02"
	defaultRollingMax = "7"
)

var once sync.Once

type logicLoggerBuilder struct {
	jobId   string
	traceId string
	userId  string
}

func initLogicLogger() {

	// SetReportCaller sets whether the standard logger will include the calling
	logrus.SetReportCaller(true)

	// Set logic logger formatter
	// 根据项目环境变量设置时间格式（默认time.RFC3339）
	timeFormat := os.Getenv(env.LogTimeFormat)
	switch timeFormat {
	case "":
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat: defaultTimestampFormat})
	default:
		logrus.SetFormatter(&LogicLoggerFormatter{TimestampFormat: timeFormat})
	}

	//根据项目环境变量设置翻滚时间格式（默认每日一翻滚）
	logRollTimeFormat := os.Getenv(env.LogRollingTimeFormat)
	if logRollTimeFormat == "" {
		logRollTimeFormat = defaultRollingTimeFormat
	}

	//根据项目环境变量设置日志文件保存的最大数量（默认7）
	logRollingMax := os.Getenv(env.LogRollingMax)
	if logRollingMax == "" {
		logRollingMax = defaultRollingMax
	}
	maxRoll, err := strconv.Atoi(logRollingMax)
	if err != nil {
		panic(err.Error())
	}

	//根据项目环境变量设置打印还是写文件(默认打印)
	logOutput := os.Getenv(env.LogOutput)
	switch logOutput {
	case "", "console":
		logrus.SetOutput(os.Stdout)
	default:
		hook, err := rollingfile.NewRollingFileTimeHook(logOutput, logRollTimeFormat, maxRoll)
		if err != nil {
			panic(err)
		}
		logrus.AddHook(hook)
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
	once.Do(initLogicLogger)
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
//func (lg *logicLoggerBuilder) AddHook(hook logrus.Hook) {
//	logrus.AddHook(hook)
//}
