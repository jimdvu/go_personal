package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
  "html/template"
  "regexp"
  "errors"
	"time"
)

type Index struct {
	Name string
	Time string
}



func main() {


	index := Index{"Jimmy", time.Now()}
 	templates = template.Must(template.ParseFiles("index.html"))
	fmt.Println("Listening");
	fmt.Println(time.Now());

	//Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files.

	http.Handle("/static/", //final url can be anything
		 http.StripPrefix("/static/",
				http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
				//url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
				//once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
				//It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

		 //Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		 if name := r.FormValue("name"); name != "" {
				index.Name = name;
		 }
		 //If errors show an internal server error message
		 //I also pass the welcome struct to the welcome-template.html file.
		 if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		 }
	})


	log.Fatal(http.ListenAndServe("", nil))
}
