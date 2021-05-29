package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.l.Println("Something went wrong")

		http.Error(rw, "An error occured", http.StatusBadGateway)
	}

	fmt.Fprintf(rw, "Hello %s", b)
}
