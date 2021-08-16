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

type Page struct {
	Title string
	Body  []byte
	Time string
}

var templates = template.Must(template.ParseFiles())


func main() {
	welcome := Welcome{time.Now().Format(time.Stamp)}

	log.Fatal(http.ListenAndServe("", nil))
}
