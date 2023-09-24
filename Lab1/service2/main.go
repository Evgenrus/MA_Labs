package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/service2/getString", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service2")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8084", nil))
}
