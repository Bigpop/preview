package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Url struct {
    UrlPath string
}

func previewFunc(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path[1:]
	fmt.Println(url)
	t := template.Must(template.ParseFiles("template.html"))
	data := Url{
		UrlPath: url,
	}
	t.Execute(w, data)
}

func templateHandler(w http.ResponseWriter, r *http.Request) { 
      w.Header().Set("Content-Type", "text/html; charset=utf-8") 
      t, err := template.ParseFiles("preview.html") 
      if err != nil { 
             fmt.Fprintf(w, "Unable to load template") 
        } 
     t.Execute(w, nil) 
}
func main() {
	http.HandleFunc("/", templateHandler)
    http.ListenAndServe(":8090", nil)
}
