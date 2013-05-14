package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("_site")))
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
