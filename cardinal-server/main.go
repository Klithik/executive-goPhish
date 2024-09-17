package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const filename string = "appkey.txt"

func generatepassword(length int) {
    log.Println("[LOG] Checking for password")
    _, err := os.Stat(filename)
    if err != nil {
        log.Println("[LOG] Password file found, importing...")
        dat, err := os.ReadFile(filename)
        if err != nil {
            panic(err)
        }
        log.Println("[LOG] Encrypted password: " + string(dat))
    }
    log.Println("[LOG] Password file not found, creating new password")

    const charset = "abcdefghijklmnopqrstuvwxyz0123456789$"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
    fmt.Println("App password: " + string(b))

    //Encrypts password and checks for its success 
    bytes, err := bcrypt.GenerateFromPassword([]byte(string(b)), 14)
    if err != nil {
        panic(err)
    }
    pass := string(bytes)

    os.WriteFile(filename, []byte(pass), 644)
    log.Println("Password created, encrypted and saved into "+filename)
}

func comparePassword(hash string) bool {
    password,err := os.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    // THIS FUNCTION ALWAYS RETURNS ITS ERROR, SO IT MUST BE NIL IF THE PASSWORD
    // IS CORRECT
    if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
        return false
    }
    return true
}

func main(){
    mux := http.NewServeMux()
    log.Print("[LOG] Listening...")

    go mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("Funciona")
    })

    http.ListenAndServe(":5000",mux)
}
