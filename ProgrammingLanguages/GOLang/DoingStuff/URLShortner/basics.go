package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

// URL mapping ke liye map
var urlMap = make(map[string]string)

func generateShortCode() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano()) // Randomness better karne ke liye
    b := make([]byte, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortCode := r.URL.Path[1:]

    fmt.Println("Redirect request received for:", shortCode)

    if longURL, exists := urlMap[shortCode]; exists {
        fmt.Println("Redirecting to:", longURL)
        http.Redirect(w, r, longURL, http.StatusMovedPermanently)
    } else {
        fmt.Println("Short URL not found:", shortCode) 
        http.Error(w, "Short URL not found", http.StatusNotFound)
    }
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    longURL := r.FormValue("url") // Assume a form with "url" field
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    shortCode := generateShortCode()
    urlMap[shortCode] = longURL

    fmt.Println("Stored:", shortCode, "=>", longURL)

    fmt.Fprintf(w, "Short URL: http://localhost:8080/%s", shortCode)
}

func main() {
    http.HandleFunc("/", redirectHandler)
    http.HandleFunc("/shorten", createShortURL)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}