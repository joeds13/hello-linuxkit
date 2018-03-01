package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/linuxkit", helloLinuxKit)

	log.Println("Listening on 0.0.0.0:1337")

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Saying \"hello\" to the world!")
	fmt.Fprintf(w, "Hello, world!")
}

func helloLinuxKit(w http.ResponseWriter, r *http.Request) {
	log.Println("Saying \"hello\" to linuxkit!")
	fmt.Fprintf(w, "Hello, LinuxKit!")
}
