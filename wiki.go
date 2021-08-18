package main

import (
	"io/ioutil"
	"log"
	"net/http"
  "html/template"
  "regexp"
  "errors"
	"time"
	"fmt"
)

type Index struct {
	Name string
	Time string
}



func main() {


	index := Index{"Jimmy", time.Now().Format(time.Stamp)}
 	templates = template.Must(template.ParseFiles("templates/index.html"))



	log.Fatal(http.ListenAndServe("", nil))
}
