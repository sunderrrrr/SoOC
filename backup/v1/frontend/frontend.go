package frontend

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../html/index.html")
		fmt.Println("Запрос на Index")
		return
	}
	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "../html/style.css")
		return
	}
	if r.URL.Path == "/parse.js" {
		http.ServeFile(w, r, "../html/parse.js")
		return
	}
}

// Чтобы раздавать файл нужно использовать путь /эндпоинт/нужный файл
func Create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/create" {
		http.ServeFile(w, r, "../html/create.html")
		return
	}

	if r.URL.Path == "/create" {
		http.ServeFile(w, r, "../html/style.css")
		return
	}
	if r.URL.Path == "/create" {
		http.ServeFile(w, r, "../html/parse.js")
		return
	}
}

func Guide(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/guide" {
		http.ServeFile(w, r, "../html/guide.html")
		return
	}

	if r.URL.Path == "/guide/style.css" {
		http.ServeFile(w, r, "../html/style.css")
		return
	}

}
