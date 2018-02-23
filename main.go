package main

import (
	"log"
	"net/http"

	bs "github.com/collins-b/basefiles"
)

func main() {
	log.Println("Server running at: http://localhost:3000")
	http.HandleFunc("/", bs.Index)
	http.HandleFunc("/show", bs.Show)
	http.HandleFunc("/new", bs.New)
	http.HandleFunc("/edit", bs.Edit)
	http.HandleFunc("/insert", bs.Insert)
	http.HandleFunc("/update", bs.Update)
	http.HandleFunc("/delete", bs.Delete)
	http.ListenAndServe(":3000", nil)
}
