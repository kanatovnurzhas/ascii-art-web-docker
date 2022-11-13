package delivery

import (
	"net/http"
	"text/template"
)

func HandleErr(w http.ResponseWriter, status int) {
	tmp, _ := template.ParseFiles("templates/error.html")
	data := Data{
		ErrStatus: status,
		Err:       http.StatusText(status),
	}
	w.WriteHeader(status)
	tmp.Execute(w, data)
}
