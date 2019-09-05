package log

import (
	"fmt"
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

// NewLogicLogger => Generate a LogicLogger.
// Require env $PROJECT_NAME $BP_LOG_TIME_FORMAT $BP_LOG_OUTPUT $BP_LOG_LEVEL
func NewLogicLogger() (lg *logicLogger) {
	initLogicLogger()
	return new(logicLogger)
}

func newLogicLogger(lg *logicLogger) *logrus.Entry {
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

func (lg *logicLogger) Trace(args ...interface{}) {
	newLogicLogger(lg).Trace(args...)
}

func (lg *logicLogger) Debug(args ...interface{}) {
	newLogicLogger(lg).Debug(args...)
}

func (lg *logicLogger) Print(args ...interface{}) {
	newLogicLogger(lg).Print(args...)
}

func (lg *logicLogger) Info(args ...interface{}) {
	newLogicLogger(lg).Info(args...)
}

func (lg *logicLogger) Warn(args ...interface{}) {
	newLogicLogger(lg).Warn(args...)
}

func (lg *logicLogger) Warning(args ...interface{}) {
	newLogicLogger(lg).Warning(args...)
}

func (lg *logicLogger) Error(args ...interface{}) {
	newLogicLogger(lg).Error(args...)
}

func (lg *logicLogger) Fatal(args ...interface{}) {
	newLogicLogger(lg).Fatal(args...)
}

func (lg *logicLogger) Panic(args ...interface{}) {
	newLogicLogger(lg).Panic(args...)
	panic(fmt.Sprint(args...))
}

func (lg *logicLogger) Tracef(format string, args ...interface{}) {
	newLogicLogger(lg).Tracef(format, args...)
}

func (lg *logicLogger) Debugf(format string, args ...interface{}) {
	newLogicLogger(lg).Debugf(format, args...)
}

func (lg *logicLogger) Infof(format string, args ...interface{}) {
	newLogicLogger(lg).Infof(format, args...)
}

func (lg *logicLogger) Printf(format string, args ...interface{}) {
	newLogicLogger(lg).Printf(format, args...)
}

func (lg *logicLogger) Warnf(format string, args ...interface{}) {
	newLogicLogger(lg).Warnf(format, args...)
}

func (lg *logicLogger) Warningf(format string, args ...interface{}) {
	newLogicLogger(lg).Warningf(format, args...)
}

func (lg *logicLogger) Errorf(format string, args ...interface{}) {
	newLogicLogger(lg).Errorf(format, args...)
}

func (lg *logicLogger) Fatalf(format string, args ...interface{}) {
	newLogicLogger(lg).Fatalf(format, args...)

}

func (lg *logicLogger) Panicf(format string, args ...interface{}) {
	newLogicLogger(lg).Panicf(format, args...)
}

func (lg *logicLogger) Traceln(args ...interface{}) {
	newLogicLogger(lg).Traceln(args...)
}

func (lg *logicLogger) Debugln(args ...interface{}) {
	newLogicLogger(lg).Debugln(args...)
}

func (lg *logicLogger) Infoln(args ...interface{}) {
	newLogicLogger(lg).Infoln(args...)
}

func (lg *logicLogger) Println(args ...interface{}) {
	newLogicLogger(lg).Println(args...)
}

func (lg *logicLogger) Warnln(args ...interface{}) {
	newLogicLogger(lg).Warnln(args...)
}

func (lg *logicLogger) Warningln(args ...interface{}) {
	newLogicLogger(lg).Warningln(args...)
}

func (lg *logicLogger) Errorln(args ...interface{}) {
	newLogicLogger(lg).Errorln(args...)
}

func (lg *logicLogger) Fatalln(args ...interface{}) {
	newLogicLogger(lg).Fatalln(args...)

}

func (lg *logicLogger) Panicln(args ...interface{}) {
	newLogicLogger(lg).Panicln(args...)
}
