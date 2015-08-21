package headerhandler_test

import (
	"github.com/myshkin5/webert/headerhandler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("HeaderHandler", func() {
	It("adds headers then allows its inner handler to continue processing", func() {
		headers := make(map[string][]string)
		headers["first"] = []string{"first-value"}
		headers["second"] = []string{"second-value", "second-second-value"}

		innerHandler := mockHandler{
			responses: make(chan http.ResponseWriter, 10),
			requests:  make(chan *http.Request, 10),
		}

		handler := headerhandler.New(&innerHandler, headers)

		response := httptest.NewRecorder()
		request := &http.Request{
			RequestURI: "/index.html",
		}

		handler.ServeHTTP(response, request)

		Expect(innerHandler.responses).To(Receive(Equal(response)))
		Expect(innerHandler.requests).To(Receive(Equal(request)))
		Expect(response.Header().Get("first")).To(Equal("first-value"))
		Expect(response.Header().Get("second")).To(Equal("second-value"))
	})
})

type mockHandler struct {
	responses chan http.ResponseWriter
	requests  chan *http.Request
}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.responses <- w
	m.requests <- r
}
