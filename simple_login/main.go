package main

import (
        "log"
        "net/http"
)

func main() {
        // Serve index.html from the same directory
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "index.html")
        })

        // Start the server on port 8080
        log.Println("Server listening on port 8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatalf("ListenAndServe: %v", err)
        }
}
