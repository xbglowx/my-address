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

type httpResponseWriter struct {
	http.ResponseWriter
}

func (x *httpResponseWriter ) logClientIp(msg string) {
	io.WriteString(x.ResponseWriter, msg)
	log.Println(msg)
}

func getRequestIpHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"x-forwarded-for", "true-client-ip"}
	msg := fmt.Sprintf("Looking for headers %v\n", headers)

	log := httpResponseWriter{w}
	log.logClientIp(msg)

	for _, header := range headers {
		clientIp := r.Header.Get(header)
		msg = fmt.Sprintf("%v: %v\n", header, clientIp)
		log.logClientIp(msg)
	}
	msg = fmt.Sprintf("Remote Addr: %v\n", r.RemoteAddr)
	log.logClientIp(msg)
}