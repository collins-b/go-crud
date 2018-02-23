package basefiles

import (
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	lId := r.URL.Query().Get("id")
	d, err := db.Query("SELECT * FROM langs WHERE id=?", lId)
	if err != nil {
		panic(err.Error())
	}
	lang := ProgrammingLanguage{}
	for d.Next() {
		var id int
		var name, typlel string
		err = d.Scan(&id, &name, &typlel)
		if err != nil {
			panic(err.Error())
		}
		lang.Id = id
		lang.Name = name
		lang.TypeL = typlel
	}
	tmpl.ExecuteTemplate(w, "Show", lang)
}
