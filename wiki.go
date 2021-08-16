package main

import (
	"io/ioutil"
	"log"
	"net/http"
  "html/template"
  "regexp"
  "errors"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles())


func main() {

	log.Fatal(http.ListenAndServe("", nil))
}
