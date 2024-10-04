package main

import (
	"article-hub/internal"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

var container = internal.NewContainer()

func addArticleHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Printf("Method \"%s\" Not Allowed", request.Method)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "Error reading body", http.StatusInternalServerError)
		return
	}
	article := internal.Article{}
	err = json.Unmarshal(data, &article)
	if err != nil {
		log.Printf("Error parsing body: %v", err)
		http.Error(writer, "Error parsing body", http.StatusInternalServerError)
		return
	}
	err = container.AddArticle(article)
	if err != nil {
		log.Printf("Error adding article: %v", err)
		http.Error(writer, "Error adding article", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Article added successfully"))
	if err != nil {
		log.Printf("Error writing body: %v", err)
		http.Error(writer, "Error writing body", http.StatusInternalServerError)
		return
	}
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		log.Printf("Method \"%s\" Not Allowed\n", request.Method)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := request.URL.Query().Get("query")
	if query == "" {
		log.Println("Query parameter is missing")
		http.Error(writer, "Missing query", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(query)
	if err != nil {
		log.Printf("Query parameter is invalid: %s\n", query)
	}
	art, err := container.Search(id)
	if err != nil {
		log.Printf("Search Error: %s\n", err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(art)
	if err != nil {
		log.Println("Failed to write create JSON object")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		log.Println("Failed to write response")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/add", addArticleHandler)
	log.Printf("Server started at http://localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
