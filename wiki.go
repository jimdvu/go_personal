package main

import (
	"io/ioutil"
	"log"
	"net/http"
  "html/template"
  "regexp"
  "errors"
	"time"
)

type Landing struct {
	Time string
}



func main() {


	index := Index{"Anonymous", time.Now().Format(time.Stamp)}

 	templates = template.Must(template.ParseFiles("landing.html"))


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
				welcome.Name = name;
		 }
		 //If errors show an internal server error message
		 //I also pass the welcome struct to the welcome-template.html file.
		 if err := templates.ExecuteTemplate(w, "landing.html", welcome); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		 }
	})

	log.Fatal(http.ListenAndServe("", nil))
}
