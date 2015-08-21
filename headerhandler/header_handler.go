package headerhandler

import "net/http"

type HeaderHandler struct {
	innerHandler http.Handler
	headers      map[string][]string
}

func New(innerHandler http.Handler, headers map[string][]string) *HeaderHandler {
	return &HeaderHandler{
		innerHandler: innerHandler,
		headers:      headers,
	}
}

func (h *HeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key, values := range h.headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	h.innerHandler.ServeHTTP(w, r)
}
