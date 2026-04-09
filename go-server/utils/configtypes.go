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
	INFO    LogLevel = "INFO"
	DEBUG   LogLevel = "DEBUG"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

// type for logging on the server
type LoggerConfig struct {
	FileOutputHandler []LogFileHandler `yaml:"log-file-handlers"`
}

// type for file handlers
type LogFileHandler struct {
	FilePath string   `yaml:"file-path"`
	LogLvl   LogLevel `yaml:"level"`
}
