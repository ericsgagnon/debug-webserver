//go:generate rm assets.go
//go:generate go run embed.go
//go:generate go run main.go assets.go

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var err error

type server struct {
	r      *mux.Router
	Assets http.FileSystem
}

func main() {

	s := server{}
	s.Assets = assets
	s.r = mux.NewRouter()
	s.routes()

	log.Fatal(http.ListenAndServe(":18080", s.r))
}

func (s *server) routes() {
	s.r.PathPrefix("/").HandlerFunc(s.index())
}

func (s *server) index() http.HandlerFunc {

	f, err := s.Assets.Open("index.gohtml")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	tbs, err := ioutil.ReadAll(f)
	ts := string(tbs)
	t, err := template.New("template").Parse(ts)
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
