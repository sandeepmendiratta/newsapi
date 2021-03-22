package main

import (
	"fmt"

	"github.com/sandeepmendiratta/newsapi/app"
	"github.com/sandeepmendiratta/newsapi/config"
	logConfig "github.com/sandeepmendiratta/newsapi/log"
	"github.com/spf13/viper"
)

var (
	Version, Build string
)

func main() {
	config.LoadConfig()
	logConfig.InitializeLogging()
	fmt.Println("Newsapi Version:", viper.GetString("version"))
	fmt.Println("Newsapi Build Version:", viper.GetString("build"))
	fmt.Println("Welcome to the newsapi app...")
	fmt.Println("Starting application... localhost:", viper.GetString("port"))
	app.StartApp()

}
