package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"chaos-io/callchain/B/handler"
)

var (
	host        = flag.String("HOST", "http://127.0.0.1", "HOST for the server")
	httpAddress = flag.String("HTTP_ADDRESS", ":40011", "HTTP address for the server")
	callAddress = flag.String("CALL_ADDRESS", ":40021", "Call address for the server")
)

func main() {
	flag.Parse()
	errc := make(chan error)

	go handler.InterruptHandler(errc)

	go func() {
		r := mux.NewRouter()
		b := handler.NewB(*host+*httpAddress, *host+*callAddress)
		r.HandleFunc("/", b.CallItself)
		r.HandleFunc("/C", b.CallC)
		r.Use(handler.MiddlewareTrace)
		//r.Use(mux.CORSMethodMiddleware(r))
		h := cors.AllowAll().Handler(r)
		errc <- http.ListenAndServe(*httpAddress, h)
	}()

	log.Print("server run error:", <-errc)
}
