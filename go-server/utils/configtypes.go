package utils

/**
*
* BRIEF: This file will decalare types for the various config types
* AUTHOR: Elias Lopes
 */

// parent object wrapper
type Config struct {
	ServerConfig ServerConfig `yaml:"dev-server-config"`
	LogConfig    LoggerConfig `yaml:"server-log-config"`
}

// type to map struct to yaml fields
type ServerConfig struct {
	ServerPort    int      `yaml:"port"`
	CorsAllowList []string `yaml:"allow-list"`
	MethodList    []string `yaml:"method-list"`
	RequiresAuth  bool     `yaml:"requires-auth"`
}

// enum for representing log levels
type LogLevel string

const (
	NO_LOG  LogLevel = "NO_LOG"
	INFO    LogLevel = "INFO"
	DEBUG   LogLevel = "DEBUG"
	WARNING LogLevel = "WARNING"
	FATAL   LogLevel = "FATAL"
)

// type for logging on the server
type LoggerConfig struct {
	LogFiles    []string `yaml:"log-files"`
	MaxLogLevel LogLevel `yaml:"max-log-level"`
	MinLogLevel LogLevel `yaml:"min-log-level"`
}
