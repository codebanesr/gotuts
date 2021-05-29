package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handling goodbye requests")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("An error occured")

		http.Error(rw, "Something went wrong", r.Response.StatusCode)
	}

	fmt.Fprintf(rw, "Goodbye %s", b)
}
