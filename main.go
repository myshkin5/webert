package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/myshkin5/webert/headerhandler"
	"log"
	"net/http"
)

func main() {
	headers := make(map[string][]string)
	headers["first"] = []string{"first-value"}
	headers["second"] = []string{"second-value", "second-second-value"}

	router := mux.NewRouter()
	router.Handle("/", headerhandler.New(&simpleHandler{}, headers))

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

type simpleHandler struct {
}

func (h *simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello webert")
}
