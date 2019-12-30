//go:generate rm assets.go
//go:generate go run embed.go
//ignore go:generate go run main.go assets.go

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var err error

type server struct {
	r       *mux.Router
	Assets  http.FileSystem
	Address string
}

func main() {
	s := server{}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Value:       ":80",
				Destination: &s.Address,
				Usage:       "string representing the address:port to listen on",
			},
		},
		Name:  "debug-webserver",
		Usage: "runs a webserver that listens on [address] and displays http request details",
		Action: func(c *cli.Context) error {
			fmt.Println("Welcome to the Thunderdome, no logging will be provided.")
			s.Assets = assets
			s.r = mux.NewRouter()
			s.routes()
			log.Fatal(http.ListenAndServe(s.Address, s.r))
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
