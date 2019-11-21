/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/SamuelTJackson/apiMapper/xmlParser"
	"html/template"
	"net/http"
)

func main() {
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"boxY": func(height int, index int) int {
			return (index+1)*height + index*5
		},
		"textY": func(boxHeight int, index int) int {
			return (index+1)*boxHeight + index*5 + boxHeight/2
		},
	}

	tmpl, _ := template.New("display.html").Funcs(funcMap).ParseFiles("templates/display.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		root, err := xmlParser.ParseURL("http://suche.transparenz.hamburg.de/dataset/stadtrad-stationen-hamburg.xml")

		fmt.Println(err)
		if err := tmpl.Execute(w, root); err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":80", nil)

	//cmd.Execute()
}
