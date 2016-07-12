package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var port = flag.Int("port", 9000, "the port on which to serve")

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	views, err := ioutil.ReadDir("web/views")
	if err != nil {
		panic(err)
	}
	for _, t := range views {
		if t.IsDir() {
			continue
		}
		partials, err := template.New(t.Name()).ParseGlob("web/views/partials/*.html")
		templates[t.Name()] = partials
		if err != nil {
			panic(err)
		}
		templates[t.Name()].ParseFiles("web/views/" + t.Name())
	}
}

func main() {
	flag.Parse()
	router := httprouter.New()
	router.GET("/", Homepage)
	router.GET("/healthy", Healthy)
	router.GET("/d/:path", Default)
	router.ServeFiles("/public/*filepath", http.Dir("web/public"))

	fmt.Printf("Starting server on port: %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
	if err != nil {
		fmt.Printf("Unable to start server: %v\n", err)
	}
}

func Homepage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	executeTemplate(w, "index.html", map[string]interface{}{"Title": "home"})
}

func Default(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateName := lowercaseFirstChar(p.ByName("path"))
	executeTemplate(w, templateName+".html", map[string]interface{}{"Title": templateName})
}

func lowercaseFirstChar(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

func executeTemplate(w http.ResponseWriter, t string, d map[string]interface{}) error {
	// d["Bootstrap"] = "/public/css/bootstrap.css"
	d["Bootstrap"] = "//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"
	return templates[t].ExecuteTemplate(w, t, d)
}

func Healthy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("good"))
}
