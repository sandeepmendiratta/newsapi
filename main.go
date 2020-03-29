package main

import (
	"fmt"
	"github.com/sandeepmendiratta/newsapi/app"
	"github.com/sandeepmendiratta/newsapi/config"
	logConfig "github.com/sandeepmendiratta/newsapi/log"
)

var (
	Version, Build string
)

func main() {
	config.LoadConfig()
	logConfig.InitializeLogging()
	fmt.Printf("Version: %s, Build:%s\n", config.Configuration.Version, config.Configuration.Build)
	fmt.Println("Welcome to the:", config.Configuration.AppName)
	fmt.Println("Starting application... localhost:", config.Configuration.Port)
	app.StartApp()

}
