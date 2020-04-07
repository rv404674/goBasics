package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa Go is nice")
	//Fprintf write to w
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web server created by Rahul")
}

func main() {

	// "/" denotes base page. Index_handler is the function you want to run
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)

	http.ListenAndServe(":8000", nil)

}
