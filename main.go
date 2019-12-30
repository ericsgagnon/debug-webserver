//go:generate go run embed.go

package main

//ignore go:generate go get github.com/rakyll/statik
//ignore go:generate statik -src=assets -p=assets -c="Package assets embeds static assets." -f
//ignore go:generate go run -tags=dev embed.go

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	// _ "assets"
	"github.com/gorilla/mux"
)

var err error

type server struct {
	r             *mux.Router
	RootDirectory http.Dir
	File          http.File
	Assets        http.FileSystem
}

// use with go gen
//

func main() {

	// the server is all
	s := server{}
	s.Assets = assets
	// s.RootDirectory = http.Dir("assets")
	// s.File, err = s.RootDirectory.Open("index.gohtml")
	// if err != nil {
	// 	log.Fatalf("Could not read file: %s", err)
	// }

	// s.Assets = http.Dir("assets")
	// err = vfsgen.Generate(s.Assets, vfsgen.Options{
	// 	Filename:     "assets.go",
	// 	PackageName:  "main",
	// 	BuildTags:    "!dev",
	// 	VariableName: "assets",
	// })
	// if err != nil {
	// 	log.Fatalf("vfsgen failed: %s", err)
	// }

	s.r = mux.NewRouter()
	s.routes()

	// Assets := http.Dir("assets")

	log.Fatal(http.ListenAndServe(":18080", s.r))
}

func (s *server) routes() {
	s.r.PathPrefix("/").HandlerFunc(s.index())
}

func (s *server) index() http.HandlerFunc {

	tbs, err := ioutil.ReadAll(s.File)
	ts := string(tbs)
	t, err := template.New("template").Parse(ts)
	// t, err := template.ParseFiles("assets/index.gohtml")
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

// // embed uses xxx to embed the given path into an http.FileSystem in a .go file
// func embed(path string) *http.FileSystem {

// }
