package app

import (
	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	"github.com/sandeepmendiratta/newsapi/config"
	"github.com/sandeepmendiratta/newsapi/controllers"
	"log"
	"net/http"
	//"os"
)

const (
	STATIC_DIR = "/assets/"
)

func StartApp() {
	port := config.Configuration.Port
		if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/health", controllers.GetHealth).Methods("GET")
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	r.HandleFunc("/search", controllers.SearchHandler).Methods("GET")
	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	r.Use(muxlogrus.NewLogger().Middleware)
	log.Fatal(http.ListenAndServe(":"+port, r))

}