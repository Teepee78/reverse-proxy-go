package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"teepee78/reverse-proxy-go/config"
	"teepee78/reverse-proxy-go/server"
	"teepee78/reverse-proxy-go/utils"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	server.ServeStatic(w, r)
	server.ServeDynamic(w, r)
}

func main() {
	flags := config.GetFlags()
	config.GetConfig(flags.Config)

	port := utils.GetPort()

	http.HandleFunc("/", handler)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Channel to listen for OS signals (e.g., Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("Proxy running on PORT:", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	fmt.Println("\nShutting down server...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Forced to shutdown:", err)
	} else {
		fmt.Println("Server exited gracefully")
	}
}
