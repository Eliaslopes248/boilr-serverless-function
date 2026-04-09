package utils

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

/**
*
* BRIEF: This file will create the REST server runtime
* AUTHOR: Elias Lopes
 */

// parent object wrapper
type Config struct {
	ServerConfig ServerConfig `yaml:"dev-server-config"`
}

// type to map struct to yaml fields
type ServerConfig struct {
	ServerPort    int      `yaml:"port"`
	LogFile       string   `yaml:"logfile"`
	CorsAllowList []string `yaml:"allow-list"`
	RequiresAuth  bool     `yaml:"requires-auth"`
}

// creates the rest servers config
func Get_server_config() (ServerConfig, error) {
	// read yaml file into []bytes
	fileBytes, err := os.ReadFile("config/dev.yml")

	// create struct for holding config info
	var parentConfig Config
	var config ServerConfig

	// handle error
	if err != nil {
		fmt.Printf("[ERROR] When reading yaml file bytes: %v", err)
		return config, err
	}

	// extract the yaml data and map to struct
	err = yaml.Unmarshal(fileBytes, &parentConfig)

	// handle error
	if err != nil {
		fmt.Printf("[ERROR] When mapping yaml to struct: %v", err)
		return config, err
	}

	// assign config
	config = parentConfig.ServerConfig

	return config, nil
}

// sets up the server logger
func configure_logger() {

}

// initializes Gin api options for the server
func create_server_options() {

}

// creates the rest server
func Create_server() (*gin.Engine, error) {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	return r, nil
}
