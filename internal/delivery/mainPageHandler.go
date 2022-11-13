package delivery

import (
	"log"
	"net/http"
	"text/template"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErr(w, 405)
		return
	}
	if r.URL.Path != "/" {
		HandleErr(w, 404)
		return
	}
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		HandleErr(w, 404)
		return
	}

	data := Data{
		Art: "",
	}

	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Println(err)
		HandleErr(w, 500)
		return
	}
}
