package main

import (
	"html/template"
	"log"
	"net/http"

	// "github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
	// uuid "github.com/satori/go.uuid"
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
	// s.r.HandleFunc(".*", s.index())
	s.r.PathPrefix("/").HandlerFunc(s.index())
}

func (s *server) index() http.HandlerFunc {
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, r.Header)

		// c, err := r.Cookie("session")
		// if err != nil {
		// 	u := uuid.NewV4()
		// 	c = &http.Cookie{
		// 		Name:  "session",
		// 		Value: u.String(),
		// 	}
		// 	http.SetCookie(w, c)
		// }

		// data := struct {
		// 	Headers   http.Header
		// 	Cookies   []*http.Cookie
		// 	Context   context.Context
		// 	Request   *http.Request
		// 	SessionID string
		// }{
		// 	Headers:   r.Header,
		// 	Cookies:   r.Cookies(),
		// 	Context:   r.Context(),
		// 	Request:   r,
		// 	SessionID: c.String(),
		// }

		// // fmt.Println(reflect.TypeOf(data))

		// dt := reflect.TypeOf(*data.Request)
		// dv := reflect.ValueOf(*data.Request)

		// // v := make([]interface{}, dv.NumField())
		// v := make(map[string]interface{})

		// for i := 0; i < dt.NumField(); i++ {
		// 	fieldName := dv.Type().Field(i).Name
		// 	if dv.Field(i).CanInterface() {
		// 		v[fieldName] = dv.Field(i).Interface()
		// 	} else {
		// 		v[fieldName] = dv.Field(i).Type()
		// 	}
		// 	fmt.Printf("%#v\t%#v\n", fieldName, v[fieldName])
		// }
		// fmt.Println(v)

		// fmt.Println(v)

		// fmt.Fprint(w, "Request:\n", r)
		// fmt.Fprint(os.Stdout, data)
		// fmt.Fprint(w, r.Cookies())
		// fmt.Fprint(w, r.Context())
		// fmt.Fprint(w, r)

		// err = t.Execute(w, data)
		err = t.Execute(w, r)
		if err != nil {
			// log.Fatalf("Template %#v failed execution using data %#v", t, nil)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
