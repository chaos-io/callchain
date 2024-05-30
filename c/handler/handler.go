package handler

import (
	"encoding/json"
	"net/http"
)

type C struct {
	Url     string
	CallUrl string
	Target  string
}

func NewC(url, callUrl string) *C {
	return &C{
		Url:     url,
		CallUrl: callUrl,
	}
}

func (c *C) CallItself(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("C")
}
