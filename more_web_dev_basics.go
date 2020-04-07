package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// you can do a multiline print as well
	fmt.Fprintf(w, `
	<h1> Hey there </h1>
	<p> Go is fast </p>
	<p> .. and simple </p>
`)

	fmt.Fprintf(w, "<h1> Hey there </h1>")
	fmt.Fprintf(w, "<p> Go is Fast and Simple </p>")
	fmt.Fprintf(w, "<p> You %s can even add %s </p>", "can", "variables")

}

//func main() {
	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8000", nil)
}
