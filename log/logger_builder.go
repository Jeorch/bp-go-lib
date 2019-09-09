// Package log is bp-go-lib's log middleware.
package log

import "github.com/sirupsen/logrus"

type LoggerBuilder interface {
	Build() *logrus.Entry
	AddHook(hook logrus.Hook)
}

type LogicLoggerBuilder interface {
	SetJobId(jobId string) *logicLoggerBuilder
	SetTraceId(jobId string) *logicLoggerBuilder
	SetUserId(jobId string) *logicLoggerBuilder
}
