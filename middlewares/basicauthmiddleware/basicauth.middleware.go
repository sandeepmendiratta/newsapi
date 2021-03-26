package basicauthmiddleware

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("unauthorizedn.\n"))
			return
		}
		handler(w, r)
	}
}


func checkUsernameAndPassword(username, password string) bool {
	return username == "user" && password == "pass"
}


func SimpleAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if viper.GetBool("simpletokenenable") && mux.CurrentRoute(r).GetName() != "health" {
				s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
				if len(s) != 2 {
					log.Info("MiddlewareSimpleToken: No Authorization token provided")
					http.Error(w, "Not Authorized", 401)
					return
				}

				if s[0] != "Bearer" {
					log.Infof("MiddlewareSimpleToken: Authorization type %s is not valid. Expected 'Bearer'", s[0])
					http.Error(w, "Not Authorized", 401)
					return
				}
				if s[1] != viper.GetString("token") {
					log.Info("MiddlewareSimpleToken: Authorization token provided does not match")
					http.Error(w, "Not Authorized", 401)
					return
				}
			}
			h.ServeHTTP(w, r)
		},
	)
}