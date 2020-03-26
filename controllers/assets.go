package controllers
import (
	"html/template"
	"net/http"
)
var tpl = template.Must(template.ParseFiles("index.html"))
// or


func Assets(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

