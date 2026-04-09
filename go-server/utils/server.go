package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

/**
*
* BRIEF: This file will create the REST server runtime
* AUTHOR: Elias Lopes
 */

// creates the rest servers config
func Get_server_config() (Config, error) {

	// check env to see what environment type we're running in
	var configFilePath string = "config/dev.yml"

	// get environment variables
	envType := os.Getenv("BOILR_SERVER_ENV_TYPE")
	strings.ToLower(envType)

	if len(envType) > 0 && (envType == "prod" || envType == "production") {
		configFilePath = "config/prod.yml"
	}

	// read yaml file into []bytes
	fileBytes, err := os.ReadFile(configFilePath)

	// create struct for holding config info
	var parentConfig Config

	// handle error
	if err != nil {
		fmt.Printf("[ERROR] When reading yaml file bytes: %v", err)
		return parentConfig, fmt.Errorf("[ERROR] When reading yaml file bytes")
	}

	// extract the yaml data and map to struct
	err = yaml.Unmarshal(fileBytes, &parentConfig)

	// handle error
	if err != nil {
		fmt.Printf("[ERROR] When mapping yaml to struct: %v", err)
		return parentConfig, fmt.Errorf("[ERROR] When mapping yaml to struct")
	}

	fmt.Printf("LOG FILES: %v", parentConfig.LogConfig.FileOutputHandler)

	return parentConfig, nil
}

// sets up the server logger
func configure_logger(config LoggerConfig) {

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
