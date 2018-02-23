package basefiles

import (
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		typel := r.FormValue("typel")
		id := r.FormValue("uid")
		send, err := db.Prepare("UPDATE langs SET name=?, typel=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		send.Exec(name, typel, id)
		log.Println("UPDATE: Name: " + name + " Type: " + typel)
	}
	http.Redirect(w, r, "/", 301)
}
