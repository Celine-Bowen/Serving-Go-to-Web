package main

import (
	"net/http"
)

func main() {
	// fmt.Println("Hello, World")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(w, "Serving Web apps")
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)

}
