package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func myHandler1(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query()["param1"][0]
	value, _ := strconv.Atoi(param1)
	out := 1337 / value
	fmt.Println(out)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query()["param1"][0]
	value := int(param1[0])
	out := 1337 / value
	fmt.Println(out)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/abc", myHandler1)
	http.HandleFunc("/def", myHandler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
