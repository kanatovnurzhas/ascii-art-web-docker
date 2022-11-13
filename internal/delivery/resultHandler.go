package delivery

import (
	"asciiArtWeb/internal/service"
	"errors"
	"log"
	"net/http"
	"text/template"
)

func SendArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleErr(w, 405)
		return
	}

	if r.URL.Path != "/art" {
		HandleErr(w, 404)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		log.Println("Empty input")
		HandleErr(w, 400)
		return
	}

	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		HandleErr(w, 404)
		return
	}

	art, err := service.GetResult(banner, text)

	if errors.Is(err, service.ErrorCharacter) {
		log.Println(err)
		HandleErr(w, 400)
		return
	} else if err != nil {
		log.Println(err)
		HandleErr(w, 500)
		return
	}

	data := Data{
		Art: art,
	}

	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Println(err)
		HandleErr(w, 500)
		return
	}
}
