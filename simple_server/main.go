package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "On route %s\n", "/")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	http.ListenAndServe(":3000", mux)
}