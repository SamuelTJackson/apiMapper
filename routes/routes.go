package routes

import (
	"fmt"
	"github.com/SamuelTJackson/apiMapper/xmlParser"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

var funcMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"boxY": func(height int, index int) int {
		return (index+1)*height + index*5
	},
	"textY": func(boxHeight int, index int) int {
		return (index+1)*boxHeight + index*5 + boxHeight/2
	},
}

func init() {
	tmpl, _ = template.New("display.html").Funcs(funcMap).ParseFiles("templates/display.html")
}
func Frontend(w http.ResponseWriter, r *http.Request) {
	root, err := xmlParser.ParseURL("http://suche.transparenz.hamburg.de/dataset/stadtrad-stationen-hamburg.xml")

	fmt.Println(err)
	if err := tmpl.Execute(w, root); err != nil {
		fmt.Println(err)
	}
}
func Api(w http.ResponseWriter, r *http.Request) {
	if id := mux.Vars(r)["id"]; id == "" {
		log.Print("id can not be empty!")
		http.Error(w, "id can not be empty!", http.StatusBadRequest)
	}

	root, err := xmlParser.ParseURL("http://suche.transparenz.hamburg.de/dataset/stadtrad-stationen-hamburg.xml")

	fmt.Println(err)
	if err := tmpl.Execute(w, root); err != nil {
		fmt.Println(err)
	}
}
