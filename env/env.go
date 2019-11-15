// Package env is bp-go-lib's environment keys definition.
package env

const (
	//Project env key
	ProjectName = "PROJECT_NAME"

	//Log env key
	LogTimeFormat        = "BP_LOG_TIME_FORMAT"
	LogOutput            = "BP_LOG_OUTPUT"
	LogLevel             = "BP_LOG_LEVEL"
	LogRollingTimeFormat = "BP_LOG_ROLLING_TIME_FORMAT"
	LogRollingMax        = "BP_LOG_ROLLING_MAX"

	//kafka env key
	KafkaConfigEnable      = "BP_KAFKA_CONFIG_ENABLE"
	KafkaConfigPath        = "BP_KAFKA_CONFIG_PATH"
	KafkaSchemaRegistryUrl = "BP_KAFKA_SCHEMA_REGISTRY_URL"
)
