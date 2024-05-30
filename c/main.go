package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"chaos-io/callchain/C/handler"
)

var (
	host        = flag.String("HOST", "http://127.0.0.1", "HOST for the server")
	httpAddress = flag.String("HTTP_ADDRESS", ":30021", "HTTP address for the server")
	callAddress = flag.String("CALL_ADDRESS", ":30031", "Call address for the server")
)

func main() {
	flag.Parse()
	errc := make(chan error)

	go handler.InterruptHandler(errc)

	go func() {
		r := mux.NewRouter()
		c := handler.NewC(*host+*httpAddress, *host+*callAddress)
		r.HandleFunc("/", c.CallItself)
		r.Use(handler.MiddlewareTrace)
		// r.Use(mux.CORSMethodMiddleware(r))
		h := cors.AllowAll().Handler(r)
		errc <- http.ListenAndServe(*httpAddress, h)
	}()

	log.Print("server run error:", <-errc)
}
