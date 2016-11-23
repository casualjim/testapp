package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/tylerb/graceful"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		b, err := httputil.DumpRequest(req, false)
		if err != nil {
			log.Printf("handle request: %v", err)
		}
		log.Printf("%s", b)
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	httpServer := &graceful.Server{Server: new(http.Server)}
	httpServer.SetKeepAlivesEnabled(true)
	httpServer.TCPKeepAlive = 3 * time.Minute
	httpServer.ListenLimit = 50
	httpServer.Timeout = 10 * time.Second
	httpServer.Handler = mux
	httpServer.Logger = graceful.DefaultLogger()

	tcpv4Listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("serving on", tcpv4Listener.Addr().String())
	if err := httpServer.Serve(tcpv4Listener); err != nil {
		log.Fatalln(err)
	}
}
