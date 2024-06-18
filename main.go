package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type NewEventRequest struct {
	FullName  string   `json:"fullName"`
	EventDate string   `json:"eventDate"`
	EventType string   `json:"eventType"`
	Details   string   `json:"details"`
	Interest  []string `json:"interest"`
}

func main() {
	mux := http.NewServeMux()

	// serve static files eg. css, js, images
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/{filePath...}", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /create", createEvent)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(resWriter http.ResponseWriter, request *http.Request) {
	ts, err := template.ParseFiles("./ui/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(resWriter, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(resWriter, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(resWriter, "Internal Server Error", 500)
	}

}

func createEvent(resWriter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(resWriter, "Error parsing form data", http.StatusBadRequest)
		return
	}

	newEventRequest := NewEventRequest{
		FullName:  request.PostForm.Get("fullName"),
		EventDate: request.PostForm.Get("eventDate"),
		EventType: request.PostForm.Get("eventType"),
		Details:   request.PostForm.Get("details"),
		Interest:  request.PostForm["interest"],
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(resWriter).Encode(newEventRequest)
	// resWriter.Write([]byte(newEventRequest.FullName + " event has been created"))
}
