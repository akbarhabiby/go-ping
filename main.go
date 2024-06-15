package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/akbarhabiby/go-ping/helpers"
	"github.com/akbarhabiby/go-ping/middlewares"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const PORT = 3000

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", ping)

	var handler http.Handler = mux

	// * Middlewares
	handler = middlewares.RateLimiter(handler)

	// * Server
	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", PORT)
	// server.Handler = handler
	server.Handler = h2c.NewHandler(handler, &http2.Server{MaxConcurrentStreams: 500, MaxReadFrameSize: 1048576})

	fmt.Printf("ping running on port %d\n", PORT)
	server.ListenAndServe()
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte((fmt.Sprintf("pong %s", helpers.GetRealIP(req)))))
}
