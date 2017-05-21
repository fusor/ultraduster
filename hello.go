package main

import (
	"fmt"
	"log"
	"net/http"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
}

func main() {
	fmt.Println("hello from ultraduster")
	http.HandleFunc("/", PostData)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
