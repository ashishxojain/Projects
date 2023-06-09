package main

import (
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != "Get"{
		http.Error(w,"method isn't supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello")
}
func formHandler(w http.ResponseWriter, r *http.Request){
	
}
func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)  //handle it as a root 
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting the server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
	
}