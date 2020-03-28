package config

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Build        string `json:"build"`
	Version      string `json:"version"`
	ApiKey       string `json: apiKey`
	LogLevel     string `json:"logLevel"`
	AppName      string `json:"appName"`
	VersionCheck bool   `json:"versionCheck"`
	Port         string `json:"port"`
}

var Configuration *Config
var configLoaded = false

func LoadConfig() {
	Configuration = GenerateConfig()
	configLoaded = true
}

func GenerateConfig() *Config {
	c := &Config{}
	generateDefaultConfig(c)
	generateFileConfig(c)
	generateEnvConfig(c)
	generateFlagConfig(c)
	// set logLevel to overide
	enforceConfig(c)

	// log.Debug("Starting config: %+v", c)

	return c
}

//load default configuration
func generateDefaultConfig(c *Config) {
	c.Build = "test"
	c.Version = "1"
	c.LogLevel = ""
	c.AppName = "newsapi"
	c.LogLevel = ""
	c.Port = "8081"
}

//load config from file
func generateFileConfig(c *Config) {
	viper.SetConfigFile("./config.json")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s\n", err)
		log.Panicln("could not read the file")
	}
	// Confirm which config file is used
	//fmt.Println("Using config file:", viper.ConfigFileUsed())

	err := viper.Unmarshal(c)
	if err != nil {
		log.Info("could not unmarshall %v:", err)

	}
}

//load config from env
func generateEnvConfig(c *Config) {
	if value, exists := os.LookupEnv("ApiKey"); exists {
		c.ApiKey = value
	}

}

//load config from flag if required
func generateFlagConfig(c *Config) {
	flag.StringVar(&c.AppName, "appName", c.AppName, "Set the Appname")
	flag.Parse()
}

//enforce Config function is to have config variable set otherwise panics
func enforceConfig(c *Config) {
	if c.VersionCheck {
		return
	}
	//if c.ApiKey == "" {
	//	log.Panic("APiKey is required to run this - %s", c.ApiKey)
	//}
}
