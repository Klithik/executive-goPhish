package main

import (
    "fmt"
    "log"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "time"
    "math/rand"
)

func generateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func encryptPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func comparePassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main(){
    mux := http.NewServeMux()
    pass := generateRandomString(rand.Intn(5) + 12)
    fmt.Println("App password: " + pass)
    pass,err := encryptPassword(pass)
    if err != nil {
        panic(err)
    }
    log.Println("Encrypted password: " + pass)
    log.Print("Listening...")

    go mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("Funciona")
    })

    http.ListenAndServe(":5000",mux)
}
