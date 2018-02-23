package basefiles

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type ProgrammingLanguage struct {
	Id    int
	Name  string
	TypeL string
}

var db *sql.DB

func init() {
	db = conn()
}

func conn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "pl"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("forms/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	d, err := db.Query("SELECT * FROM langs ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	lang := ProgrammingLanguage{}
	res := []ProgrammingLanguage{}
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
		res = append(res, lang)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}
