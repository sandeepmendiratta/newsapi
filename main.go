package main

import (
	"fmt"
	"github.com/sandeepmendiratta/newsapi/app"
	"log"
)

//var tpl = template.Must(template.ParseFiles("index.html"))
//var apiKey *string



//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	tpl.Execute(w, nil)
//}

// these will be injected during build in build.sh script to allow printing
var (
	Version, Build string
)

func main() {
	log.Println("newsapi version:", Version, "build:", Build)
	fmt.Println("Welcome to the newsapi app...")
	fmt.Println("Starting application... localhost:8080" )
	app.StartApp()

}