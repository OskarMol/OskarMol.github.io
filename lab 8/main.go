package main

import (
	"html/template"
	"log"
	"net/http"
)

func searchPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		data := map[string]interface{}{"name": name, "text": "This is games that best suit your request: "}
		tmpl, _ := template.ParseFiles("static/search.html")
		tmpl.Execute(w, data)
		return
	}
	tmpl, _ := template.ParseFiles("static/search.html")
	tmpl.Execute(w, nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/search", searchPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
