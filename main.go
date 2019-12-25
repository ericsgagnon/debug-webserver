package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	// "github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
)

type server struct {
	r *mux.Router
}

func main() {

	s := server{}
	s.r = mux.NewRouter()
	s.routes()

	testCrap()

	log.Fatal(http.ListenAndServe(":18080", s.r))
}

func (s *server) routes() {
	s.r.HandleFunc("/", s.index())
}

func (s *server) index() http.HandlerFunc {
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, r.Header)

		data := struct {
			Headers http.Header
			test    string
		}{
			Headers: r.Header,
			test:    "stupid",
		}

		fmt.Fprint(os.Stdout, data)
		// fmt.Fprint(w, r.Cookies())
		// fmt.Fprint(w, r.Context())
		// fmt.Fprint(w, r)

		err = t.Execute(w, nil)
		if err != nil {
			// log.Fatalf("Template %#v failed execution using data %#v", t, nil)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func testCrap() {
	fmt.Fprintf(os.Stdout, "Test")
}
