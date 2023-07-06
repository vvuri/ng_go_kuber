package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID       int
	Title    string
	Authors  string
	ImageURL string
	Comment  string
	Rating   byte
}

var serverPort string = "8000"
var books []Book

func booksList(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(books)
	case "POST":
		//body, _ := ioutil.ReadAll(r.Body)
		//var book Book
		//json.Unmarshal(body, &books)
		//books = append(books, book)
		json.NewEncoder(w).Encode(books)
	}
}

func startServer() {
	http.HandleFunc("/api/list", booksList)
	// http.HandleFunc("/api/", )
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	// serverPort = os.Getenv("SERVERPORT")
	books = append(books, Book{ID: 1, Title: "Book_1"})
	log.Println("Server started on port :" + serverPort)
	startServer()
	// wget -qO- http://localhost:8000/api/list
}
