package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type A struct {
	Url     string
	CallUrl string
	Target  string
}

func NewA(url, callUrl string) *A {
	return &A{
		Url:     url,
		CallUrl: callUrl,
	}
}

func (a *A) CallItself(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("a")
}

func (a *A) CallB(w http.ResponseWriter, r *http.Request) {
	if err := httpGet(a.CallUrl, &a.Target); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%q\n", a.Target)
}

func (a *A) CallC(w http.ResponseWriter, r *http.Request) {
	if err := httpGet(a.CallUrl+"/c", &a.Target); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%q\n", a.Target)
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
