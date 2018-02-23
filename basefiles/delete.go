package basefiles

import (
	"log"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("id")
	del, err := db.Prepare("DELETE FROM langs WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(lang)
	log.Println("Language deleted. ID: " + lang)
	http.Redirect(w, r, "/", 301)
}
