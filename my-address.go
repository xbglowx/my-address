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
	headers := []string{"x-forwarded-for", "true-client-ip"}
	msg := fmt.Sprintf("Looking for headers %v\n", headers)
	io.WriteString(w, msg)
	log.Println(msg)

	for _, header := range headers {
		clientIp := r.Header.Get(header)
		msg = fmt.Sprintf("%v: %v\n", header, clientIp)
		io.WriteString(w, msg)
		log.Println(msg)
	}
	msg = fmt.Sprintf("Remote Addr: %v\n", r.RemoteAddr)
	io.WriteString(w, msg)
	log.Println(msg)
}