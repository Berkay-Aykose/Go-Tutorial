package main

import (
	"net/http"
	"text/template"
)

func main() {
	// "/assets/" yolundaki talepleri ./kaynaklar dizinindeki dosyalara yönlendirir
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./kaynaklar"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}
