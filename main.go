package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()
	// serve static files eg. css, js, images
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/{filePath...}", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /create", createTask)

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

func createTask(resWriter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		resWriter.Write([]byte("Invalid request"))
		return
	}

	fmt.Println(request.PostForm)
	resWriter.Write([]byte("Posting titles!"))
	// err := request.ParseForm()
	// if err != nil {
	// 	resWriter.WriteHeader(http.StatusBadRequest)
	// 	resWriter.Write([]byte("Invalid request"))
	// 	return
	// }
	// resWriter.Write([]byte(request.PostForm.Get("title")))
}
