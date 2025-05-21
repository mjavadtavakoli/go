package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func benchmarkmain() {

	go func() {
		fmt.Println(http.ListenAndServe("localhost:1725", nil))
	}()

	var users int

	fmt.Println("please enter number : ")
	fmt.Scanln(&users)

}
