package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template
var mesaj string = "Go ile veri gönderme"

func init() {
	tmpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index fonksiyonu, HTTP isteği alındığında şablona mesaj değişkenini gönderir ve şablonu yanıt olarak döner
func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", mesaj) //"index.html" şablonunu mesaj değişkeni ile birlikte yürütür
}
