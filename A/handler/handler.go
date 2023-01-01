package handler

import (
	"fmt"
	"net/http"
)

func CallA(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "A")
}

func CallB(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "B")
}

func CallC(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "C")
}


