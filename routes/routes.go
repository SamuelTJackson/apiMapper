package routes

import (
	"fmt"
	"github.com/SamuelTJackson/apiMapper/db"
	"github.com/SamuelTJackson/apiMapper/xmlParser"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
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
	id := mux.Vars(r)["id"]
	url, err := db.GetURLForID(id)
	log.Println(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := http.NewRequest(r.Method, url, r.Body)
	request.Header = r.Header
	log.Println(r.Method)
	client := http.Client{}
	log.Println(r.URL.Host)
	resp, err := client.Do(request)
	if err != nil {
		log.Println(resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

}
