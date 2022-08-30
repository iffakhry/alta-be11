package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

var data = []Article{
	{1, "Judul 1", "Deskripsi 1"},
	{2, "Judul 2", "Deskripsi 2"},
	{3, "Judul 3", "Deskripsi 3"},
}

func Articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		fmt.Println("request GET incoming")
		result, err := json.Marshal(data)
		if err != nil {
			log.Fatal("error marshal json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		fmt.Println("request POST incoming")
		return
	}
	http.Error(w, "Error", http.StatusBadRequest)

}

func main() {
	http.HandleFunc("/articles", Articles)
	http.HandleFunc("/books", Articles)
	fmt.Println("run server")
	http.ListenAndServe(":80", nil)
}
