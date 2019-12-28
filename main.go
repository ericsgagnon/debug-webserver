package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	r *mux.Router
}

func main() {

	s := server{}
	s.r = mux.NewRouter()
	s.routes()

	log.Fatal(http.ListenAndServe(":18080", s.r))
}

func (s *server) routes() {
	s.r.PathPrefix("/").HandlerFunc(s.index())
}

func (s *server) index() http.HandlerFunc {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
