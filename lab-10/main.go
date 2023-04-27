package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type Client struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
}

func (s *server) allPage(w http.ResponseWriter, r *http.Request) {
	var clients []Client
	res, _ := s.db.Query("select * from clients;")
	for res.Next() {
		var client Client
		err := res.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Phone)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, client)
	}

	t, _ := template.ParseFiles("static/all.html")
	t.Execute(w, clients)
}

func (s *server) createPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		phone := r.FormValue("phone")
		_, err := s.db.Exec("insert into clients(First_name, Last_name, phone) values ($1, $2, $3)", first_name, last_name, phone)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	t, _ := template.ParseFiles("static/create.html")
	t.Execute(w, nil)
}

func (s *server) deletePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		if _, err := s.db.Exec("delete from clients where user_id=$1", id); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	t, _ := template.ParseFiles("static/delete.html")
	t.Execute(w, nil)
}

func (s *server) updatePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		phone := r.FormValue("phone")
		if _, err := s.db.Exec("update clients set First_name=$1, Last_name=$2, phone=$3 where user_id=$4", first_name, last_name, phone, id); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	t, _ := template.ParseFiles("static/update.html")
	t.Execute(w, nil)
}

type server struct {
	db *sql.DB
}

func dbConnect() server {
	database, _ := sql.Open("sqlite3", "database.db")
	server := server{db: database}
	return server
}

func main() {
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/all", s.allPage)
	http.HandleFunc("/create", s.createPage)
	http.HandleFunc("/update", s.updatePage)
	http.HandleFunc("/delete", s.deletePage)

	http.ListenAndServe(":8080", nil)
}
