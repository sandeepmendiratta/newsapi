package config

import (
	"github.com/spf13/pflag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)



func LoadConfig() {
	LoadDefaultConfig()
	LoadEnvironmentConfig()
	LoadFileConfig()
	LoadFlagConfig()
	//CheckRequiredConfig()
}

func LoadDefaultConfig() {
	viper.SetDefault(LOGLEVEL, "info")
	viper.SetDefault("port", ":8081")
	viper.SetDefault("build", "test")
	viper.SetDefault("appname", "newsapi")
	viper.SetDefault("version", "1")
	viper.SetDefault("disableauth", true)
	viper.SetDefault("token", "temptoken")
	viper.SetDefault("simpletokenenable", true)

}

func LoadEnvironmentConfig() {
	viper.AutomaticEnv()
}

func LoadFileConfig() {
	if _, err := os.Stat("./config.json"); os.IsNotExist(err) {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigName("config.json")
	}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("Unable to read in config file. %v\n", err)
	}
}


func LoadFlagConfig() {
	pflag.String("ApiKey", "", "ApiKey from flag.")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func CheckRequiredConfig() {
	required := []string{APIKEY}
	for _, v := range required {
		if viper.GetString(v) == "" {
			log.Fatalf("Required config '%v' is not defined.\n", v)
		}
	}
}
