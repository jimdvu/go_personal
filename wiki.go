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


	index := Index{"Jimmy", time.Now().Format("15:04:05")}
 	templates = template.Must(template.ParseFiles("index.html"))



	log.Fatal(http.ListenAndServe("", nil))
}
