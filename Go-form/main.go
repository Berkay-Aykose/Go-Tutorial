package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

// init fonksiyonu, program çalışmadan önce bir kez çalıştırılır ve şablonları yükler
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*")) // templates klasöründeki tüm şablonları yükler
}

func main() {
	// formda action="/login" URL'sine gelen istekleri LoginHandler fonksiyonuna yönlendirir
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Eğer gelen istek "GET" metodu değilse, kullanıcıyı ana sayfaya yönlendirir
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// "Kullanici_adi" ve "Kullanici_sifre" isimli form alanlarından gelen değerleri alır
	Kullanici_adi := r.FormValue("Kullanici_adi") //name içine yazığımız yerden aldık
	Sifre := r.FormValue("Kullanici_sifre")

	if (Kullanici_adi == "Berkay") && (Sifre == "123") {
		// "Login.html" şablonunu kullanarak yanıtı oluşturur ve kullanici_adi değişkenini şablona aktarır
		tmpl.ExecuteTemplate(w, "Login.html", Kullanici_adi)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
