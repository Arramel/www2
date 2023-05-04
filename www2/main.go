package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() error {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("C:\\Users\\User\\GolandProjects\\www2\\static"))))
	http.HandleFunc("/", index)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	err := handleFunc()
	if err != nil {
		log.Fatal(err)
	}
}
