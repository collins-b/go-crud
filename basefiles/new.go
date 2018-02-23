package basefiles

import (
	"log"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		typel := r.FormValue("typel")
		send, err := db.Prepare("INSERT INTO langs(name, typel) VALUES(?, ?)")
		if err != nil {
			panic(err.Error())
		}
		send.Exec(name, typel)
		log.Println("INSERT: Name: " + name + "| Type: " + typel)
	}
	http.Redirect(w, r, "/", 301)
}
