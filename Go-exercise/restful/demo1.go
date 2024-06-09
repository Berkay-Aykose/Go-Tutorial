package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ToDo yapısı, JSON verilerini temsil eder
type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Demo1 fonksiyonu, bir HTTP GET isteği yapar ve yanıtı işler
func Demo1() {
	// Belirtilen URL'ye HTTP GET isteği yapar
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		// Eğer bir hata oluşursa, hata mesajını yazdırır
		fmt.Println(err)
	}
	// Fonksiyon sonunda HTTP yanıt gövdesini kapatır
	defer response.Body.Close()

	// HTTP yanıt gövdesini okur ve byte dizisine atar
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// byte dizisini stringe dönüştürür
	bodyString := string(bodyBytes)
	// JSON yanıtının ham halini yazdırır
	fmt.Println(bodyString)

	// JSON verisini ToDo yapısına dönüştürmek için bir değişken oluşturur
	var gecici_todo ToDo
	// JSON verisinin kodunu çözer ve ToDo yapısına atar
	json.Unmarshal(bodyBytes, &gecici_todo)
	// ToDo yapısının içeriğini yazdırır
	fmt.Println(gecici_todo)
}

func Demo2() {
	todo := ToDo{1, 2, "Alışveriş yapılacak", false}

	jsonTodo, err := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	/**************************************/
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var gecici_todo ToDo
	json.Unmarshal(bodyBytes, &gecici_todo)
	fmt.Println(gecici_todo)
}
