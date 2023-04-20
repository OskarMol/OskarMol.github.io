package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

func clientPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/client.html")
	t.Execute(w, nil)
}

func (s *server) finalPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		phone := r.FormValue("phone")
		_, err := s.db.Exec("insert into clients(First_name, Last_name, phone) values ($1, $2, $3)", first_name, last_name, phone)
		if err != nil {
			log.Fatal(err)
		}
		data := map[string]interface{}{"first_name": first_name, "last_name": last_name, "phone": phone}
		t, _ := template.ParseFiles("static/result.html")
		t.Execute(w, data)
		return

	}
}

type server struct {
	db *sql.DB
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "database.db")
	fmt.Println("Opening database")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}

	return s
}

func main() {
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/client", clientPage)
	http.HandleFunc("/result", s.finalPage)

	http.ListenAndServe(":8080", nil)
}
