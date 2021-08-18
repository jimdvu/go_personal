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


	index := Index{time.Now().Format(time.Stamp)}

 	templates = template.Must(template.ParseFiles("landing.html"))




	log.Fatal(http.ListenAndServe("", nil))
}
