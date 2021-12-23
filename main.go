package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net"
	"os"
)

func templateHandler(w http.ResponseWriter, r *http.Request) { 
      w.Header().Set("Content-Type", "text/html; charset=utf-8") 
      t, err := template.ParseFiles("preview.html") 
      if err != nil { 
             fmt.Fprintf(w, "Unable to load template") 
        } 
     t.Execute(w, nil) 
}
func main() {
	port := os.Args[1]
 	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("http://"+ipnet.IP.String()+":"+port)
			}

		}
	}
	http.HandleFunc("/", templateHandler)
    http.ListenAndServe(":"+port, nil)
}
