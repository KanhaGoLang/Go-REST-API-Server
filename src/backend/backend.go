package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloWord(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello Sanjeev!")
}

func Run(addr string) {
	http.HandleFunc("/", helloWord)
	fmt.Println("Server started & listening on Port ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
