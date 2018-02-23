package basefiles

import (
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	lid := r.URL.Query().Get("id")
	d, err := db.Query("SELECT * FROM langs WHERE id=?", lid)
	if err != nil {
		panic(err.Error())
	}
	lang := ProgrammingLanguage{}
	for d.Next() {
		var id int
		var name, typel string
		err = d.Scan(&id, &name, &typel)
		if err != nil {
			panic(err.Error())
		}
		lang.Id = id
		lang.Name = name
		lang.TypeL = typel
	}
	tmpl.ExecuteTemplate(w, "Edit", lang)
}
