package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	addr, ok := os.LookupEnv("MY_IP_ADDR")
	if !ok {
		addr = ":8080"
	}
	log.Printf("Starting web server with address %s", addr)

	http.HandleFunc("/", getRequestIpHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Sprintf("Failed to start web server: %s", err))
	}
}

func getRequestIpHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, (fmt.Sprintf("Your address is %s", r.RemoteAddr)))
}
