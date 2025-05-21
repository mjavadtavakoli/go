package main

import (
	"fmt"
	"net/http"

	_ "bufio"
	_ "context"
	_ "encoding/json"
	_ "errors"
	_ "io"
	_ "log"
	_ "math"
	_ "os"
	_ "strconv"
	_ "strings"
	_ "time"
)

func madin() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("sesrver is running on http://localhost:1725")
	http.ListenAndServe(":1725", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello my name is mohammad javad tavakoli")
}
