package main

import (
    "fmt"
    "log"
    "net/http"
)
//this is good
func main(){
    mux := http.NewServeMux()
    log.Print("Listening...")

    go mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("Funciona")
    })

    http.ListenAndServe(":5000",mux)
}
