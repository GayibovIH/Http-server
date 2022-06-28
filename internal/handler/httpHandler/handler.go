package httpHandler

import "net/http"

type HttpHandler struct {
	fib map[int]int
}

func (hh HttpHandler) calculate(num int) {

}

func (hh HttpHandler) get(num int) int {
	return 0
}

func (hh HttpHandler) CalculateHandle(rw http.ResponseWriter, r *http.Request) {

}

func (hh HttpHandler) GetHandle(rw http.ResponseWriter, r *http.Request) {

}

func New() *HttpHandler {
	fib := make(map[int]int)
	return &HttpHandler{fib: fib}
}
