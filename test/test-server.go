//go:build !exclude

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Response structure for API responses
type Response struct {
	Message string `json:"message"`
	Port    string `json:"port"`
}

// Function to handle the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request, port string) {
	response := Response{
		Message: "Hello, world!",
		Port:    port,
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode response as JSON and send it
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// Function to handle the /goodbye route
func goodbyeHandler(w http.ResponseWriter, r *http.Request, port string) {
	response := Response{
		Message: "Goodbye, see you later!",
		Port:    port,
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode response as JSON and send it
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// Function to create and start a server on a specific port
func startServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a new multiplexer for the server
	mux := http.NewServeMux()

	// Handle routes
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloHandler(w, r, port)
	})
	mux.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		goodbyeHandler(w, r, port)
	})

	// Start the server on the specified port
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func main() {
	var wg sync.WaitGroup

	// Start two identical servers on different ports
	wg.Add(2)
	go startServer("3000", &wg) // Server 1 on port 3000
	go startServer("3001", &wg) // Server 2 on port 3001

	// Wait for both servers to finish (they won't stop unless manually terminated)
	wg.Wait()
}
