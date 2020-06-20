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
	fmt.Println("Newsapi Version:", config.Configuration.Version)
	fmt.Println("Newsapi Build Version:", config.Configuration.Build)
	fmt.Println("Welcome to the newsapi app...")
	fmt.Println("Starting application... localhost:", config.Configuration.Port)
	app.StartApp()

}
