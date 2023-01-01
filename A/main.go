package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"chaos-io/callchain/A/handler"
)

var (
	httpAddress  = flag.String("HTTP_ADDRESS", ":40001", "HTTP address for the server")
)

func main() {
	flag.Parse()
	errc := make(chan error)

	go handler.InterruptHandler(errc)

	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/", handler.CallA)
		r.HandleFunc("/B", handler.CallB)
		r.HandleFunc("/C", handler.CallC)
		r.Use(handler.MiddlewareTrace)
		h := cors.AllowAll().Handler(r)
		errc <- http.ListenAndServe(*httpAddress, h)
	}()

	log.Print("server run error:", <-errc)
}
