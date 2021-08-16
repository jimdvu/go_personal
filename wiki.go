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

type Index struct {
	Title string
	Body  []byte
	Time string
}



func main() {


	index := Index{time.Now().Format(time.Stamp)}

	var templates = template.Must(template.ParseFiles("index.html"))


	log.Fatal(http.ListenAndServe("", nil))
}
