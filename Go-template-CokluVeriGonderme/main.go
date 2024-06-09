package main

import (
	"net/http"
	"text/template"
)

var tmpl template.Template

type ogrenci struct {
	Id      int
	Adi     string
	Soyadi  string
	Telefon []string
	Notlari []Notlar
}

type Notlar struct {
	DersAdi  string
	DersNotu int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func init() {
	tmpl = *template.Must(template.ParseGlob("template/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	// örnek bir öğrenci oluşturulur
	ogr := ogrenci{
		Id:      1,
		Adi:     "berkay",
		Soyadi:  "ayköse",
		Telefon: []string{"01234567890"},
		Notlari: []Notlar{
			Notlar{DersAdi: "Matematik", DersNotu: 100},
			Notlar{DersAdi: "Fizik", DersNotu: 50},
		},
	}
	// oluşturulan öğrenci bilgisi ile "index.html" template'i doldurulur ve tarayıcıya gönderilir
	tmpl.ExecuteTemplate(w, "index.html", ogr)
}
