package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	"github.com/sandeepmendiratta/newsapi/config"
	"github.com/sandeepmendiratta/newsapi/controllers"
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
	// r.HandleFunc("/api1", controllers.GetApi1).Methods("GET")
	r.HandleFunc("/search", controllers.SearchHandler).Methods("GET")
	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	r.Handle("/api1", CheckAuthenticated(controllers.GetApi1)).Methods("GET")
	r.Use(muxlogrus.NewLogger().Middleware)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

func CheckAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// if config.Configuration.DisableAuth {
		// 	next(w, req)
		// 	return
		// }

		authorizationHeader := req.Header.Get("authorization")
		err := validateHeaderToken(authorizationHeader, config.Configuration.Token)
		if err == nil {
			next(w, req)
		} else {
			log.Println("error validating token: %v", err)

			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func validateHeaderToken(authorizationHeader string, token string) error {
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 && bearerToken[0] == "Bearer" {
			if bearerToken[1] == token {
				return nil
			} else {
				return fmt.Errorf("bearer not specified or token does not match")
			}
		}
		return fmt.Errorf("invalid header parts length or missing Bearer token type")
	} else {
		return fmt.Errorf("authorization header is blank")

	}

}
