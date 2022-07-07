package calc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func returnAB(r *http.Request) (float64, float64) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	floatA, _ := strconv.ParseFloat(a, 64)
	floatB, _ := strconv.ParseFloat(b, 64)

	return floatA, floatB
}

func returnA(r *http.Request) int {
	vars := mux.Vars(r)
	a := vars["a"]
	intA, _ := strconv.Atoi(a)

	return intA
}

func Sum(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(r)
	a += b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'E', -1, 64))
}

func Diff(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(r)
	a -= b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'E', -1, 64))
}

func Div(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(r)
	a /= b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'E', -1, 64))
}

func Mul(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(r)
	a *= b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'E', -1, 64))
}

func Fac(w http.ResponseWriter, r *http.Request) {
	a := returnA(r)
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	fmt.Fprintf(w, strconv.Itoa(result))
}
