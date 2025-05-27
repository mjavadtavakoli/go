package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go on localhost! 👋")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:24")
	err := http.ListenAndServe(":24", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
