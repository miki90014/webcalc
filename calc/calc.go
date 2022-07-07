package calc

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
)

func returnAB(w http.ResponseWriter, r *http.Request) (float64, float64) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	floatA, errA := strconv.ParseFloat(a, 64)
	floatB, errB := strconv.ParseFloat(b, 64)

	if errA != nil || errB != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Error().Err(errors.New("400")).Msg("Bad Request")
	}

	return floatA, floatB
}

func returnA(w http.ResponseWriter, r *http.Request) int {
	vars := mux.Vars(r)
	a := vars["a"]
	intA, err := strconv.Atoi(a)

	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Error().Err(errors.New("400")).Msg("Bad Request")
	}

	return intA
}

func Sum(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(w, r)
	log.Info().Str(strconv.FormatFloat(a, 'f', -1, 64), strconv.FormatFloat(b, 'f', -1, 64)).Msgf("IP: %s, URL: %s", r.Host, r.URL.Path)
	a += b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'f', -1, 64))
}

func Diff(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(w, r)
	log.Info().Str(strconv.FormatFloat(a, 'f', -1, 64), strconv.FormatFloat(b, 'f', -1, 64)).Msgf("IP: %s, URL: %s", r.Host, r.URL.Path)
	a -= b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'f', -1, 64))
}

func Div(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(w, r)
	if b == 0 {
		http.Error(w, "400 Bad Request", http.StatusNotFound)
		log.Error().Err(errors.New("400")).Msg("Bad Request")
	} else {
		log.Info().Str(strconv.FormatFloat(a, 'f', -1, 64), strconv.FormatFloat(b, 'f', -1, 64)).Msgf("IP: %s, URL: %s", r.Host, r.URL.Path)
		a /= b
		fmt.Fprintf(w, strconv.FormatFloat(a, 'f', -1, 64))
	}
}

func Mul(w http.ResponseWriter, r *http.Request) {
	a, b := returnAB(w, r)
	log.Info().Str(strconv.FormatFloat(a, 'f', -1, 64), strconv.FormatFloat(b, 'f', -1, 64)).Msgf("IP: %s, URL: %s", r.Host, r.URL.Path)
	a *= b
	fmt.Fprintf(w, strconv.FormatFloat(a, 'f', -1, 64))
}

func Fac(w http.ResponseWriter, r *http.Request) {
	a := returnA(w, r)
	log.Info().Msgf("IP: %s, URL: %s, para: %s", r.Host, r.URL.Path, strconv.Itoa(a))
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	fmt.Fprintf(w, strconv.Itoa(result))
}
