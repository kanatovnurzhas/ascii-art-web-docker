package delivery

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

type Data struct {
	Art       string
	ErrStatus int
	Err       string
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route() *http.ServeMux {
	s.mux.HandleFunc("/", Handler)
	s.mux.HandleFunc("/art", SendArt)
	s.mux.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css/"))))
	s.mux.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img/"))))
	return s.mux
}
