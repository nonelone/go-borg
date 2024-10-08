package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home%s!", r.URL.Path[1:])
}

func player_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func dm_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love DM %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/player", player_handler)
	http.HandleFunc("/dm", dm_handler)
	http.HandleFunc("/", home_handler)

	fmt.Println("Hello GöBorg!")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
