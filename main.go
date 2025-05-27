package main

import (
	"fmt"
	"net/http"
)

func main() {
	// برای سرو فایل‌های static مثل html و css
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// هندل صفحه‌ی اصلی
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("Server is running at http://localhost:24")
	err := http.ListenAndServe(":24", nil)
	if err != nil {
		fmt.Println("خطا:", err)
	}
}
