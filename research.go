package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("_site")))
	fmt.Println("listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
