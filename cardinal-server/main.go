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

func generatepassword(length int) {
    const filename string
    filename = "appkey.txt"

    log.Println("[LOG] Checking for password")
    if _, err != os.Stat(filename); os.IsNotExist(err) {
        log.Println("[LOG] Password file found, importing...")
        return nilfunc
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
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Println("[ERROR] " + string(err))
        panic(err)
    }
    pass := string(bytes)

    os.WriteFile(filename, []byte(pass), 644)
    log.Println("Password created, encrypted and saved into "+filename)
}

func encryptPassword(password string) (string, error) {
}

func comparePassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main(){
    mux := http.NewServeMux()
    log.Print("[LOG] Listening...")

    go mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("Funciona")
    })

    http.ListenAndServe(":5000",mux)
}
