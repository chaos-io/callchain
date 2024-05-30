package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type B struct {
	Url     string
	CallUrl string
	Target  string
}

func NewB(url, callUrl string) *B {
	return &B{
		Url:     url,
		CallUrl: callUrl,
	}
}

func (b *B) CallItself(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("B")
}

func (b *B) CallC(w http.ResponseWriter, r *http.Request) {
	if err := httpGet(b.CallUrl, &b.Target); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%q\n", b.Target)
}

func httpGet(url string, resp interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if resp != nil {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		if len(content) > 0 {
			if err := json.Unmarshal(content, resp); err != nil {
				return err
			}
		}
	}

	return nil
}
