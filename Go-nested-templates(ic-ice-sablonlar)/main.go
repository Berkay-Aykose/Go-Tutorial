package main

import (
	"net/http"
	"text/template"
)

// tpl değişkeni, şablonları depolamak için kullanılır
var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// init fonksiyonu, program başlatıldığında otomatik olarak çalışır ve şablonları yükler
func init() {
	// templates klasöründeki tüm dosyaları yükler ve tpl değişkenine atar
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
