package httpHandler

import (
	"net/http"
)

type HttpHandler struct {
	fib map[int]int
}

func (hh HttpHandler) calculate(num int) {
	var fibonacci = []int{0, 1}
	for i := 2; i < num; i++ {
		num2 := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, num2)
	}

	hh.fib[num] = fibonacci[num-1]
}

func (hh HttpHandler) get(num int) int {
	var value, okey = hh.fib[num]
	if okey {
		return value
	} else {
		return -1
	}
}

func (hh HttpHandler) CalculateHandle(rw http.ResponseWriter, r *http.Request) {

}

func (hh HttpHandler) GetHandle(rw http.ResponseWriter, r *http.Request) {

}

func New() *HttpHandler {
	fib := make(map[int]int)
	return &HttpHandler{fib: fib}
}
