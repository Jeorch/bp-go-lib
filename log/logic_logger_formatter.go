package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/env"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type fieldKey string

// FieldMap allows customization of the key names for default fields.
type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

type logicLoggerOutput struct {
	Time        string `json:"Time"`
	Message     string `json:"Message"`
	Hostname    string `json:"Hostname"`
	ProjectName string `json:"ProjectName"`
	File        string `json:"File"`
	Func        string `json:"Func"`
	JobId       string `json:"JobId"`
	TraceId     string `json:"TraceId"`
	UserId      string `json:"UserId"`
	Level       string `json:"Level"`
}

// LogicLoggerFormatter formats logs into parsable json
type LogicLoggerFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string
}

func (f *LogicLoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	lLogger := new(logicLoggerOutput)
	//lLogger.Time = entry.Time.Format(defaultTimestampFormat)
	lLogger.Time = entry.Time.Format(f.TimestampFormat)
	lLogger.Hostname, _ = os.Hostname()
	//根据项目环境变量设置项目名（默认空）
	lLogger.ProjectName = os.Getenv(env.ProjectName)

	funcVal := entry.Caller.Function
	fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	if funcVal != "" {
		lLogger.Func = funcVal
	}
	if fileVal != "" {
		lLogger.File = fileVal
	}

	if JobId, ok := entry.Data["JobId"]; ok {
		lLogger.JobId = JobId.(string)
	}
	if TraceId, ok := entry.Data["TraceId"]; ok {
		lLogger.TraceId = TraceId.(string)
	}
	if UserId, ok := entry.Data["UserId"]; ok {
		lLogger.UserId = UserId.(string)
	}

	lLogger.Message = entry.Message
	lLogger.Level = strings.ToUpper(entry.Level.String())

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if err := encoder.Encode(lLogger); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
