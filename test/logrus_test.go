package test

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

//func init() {
//	// Log as JSON instead of the default ASCII formatter.
//	logrus.SetFormatter(&logrus.JSONFormatter{})
//
//	// Output to stdout instead of the default stderr
//	// Can be any io.Writer, see below for File example
//	logrus.SetOutput(os.Stdout)
//}

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Trace msg")
	logrus.Debug("Debug msg")
	logrus.Info("Info msg")
	logrus.Warn("Warn msg")
	logrus.Error("Error msg")
	logrus.Fatal("Fatal msg")
	logrus.Panic("Panic msg")
}

func TestLogrusFormat(t *testing.T) {

	//AddHook function
	//errorHook := airbrake.NewHook(123, "xyz", "production")
	//errorHook.Airbrake.SetHost("http://localhost")
	//logrus.AddHook(errorHook)

	//Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:   "Time",
			logrus.FieldKeyLevel:   "Level",
			logrus.FieldKeyMsg:   "Message",
			logrus.FieldKeyFile:   "File",
			logrus.FieldKeyFunc:   "Func",
		},
	})
	//logrus.SetFormatter(new(JSONFormatter))

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)

	Hostname, _ := os.Hostname()

	ProjectName := "bp-log"
	JobId := "job-001"
	TraceId := "test-001"
	UserId := "jeorch"

	diyLogger := logrus.WithFields(logrus.Fields{"Hostname": Hostname, "ProjectName": ProjectName,
		"JobId": JobId, "TraceId": TraceId, "UserId": UserId})

	diyLogger.Trace("Trace msg")
	diyLogger.Debug("Debug msg")
	diyLogger.Info("Info msg")
	diyLogger.Warn("Warn msg")
	diyLogger.Error("Error msg")
}
