package calc

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"konta.monika/webcalc/health"

	"github.com/gorilla/mux"
)

func returnAB(w http.ResponseWriter, r *http.Request) (float64, float64, error, error) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	floatA, errA := strconv.ParseFloat(a, 64)
	floatB, errB := strconv.ParseFloat(b, 64)

	if errA != nil || errB != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Error().Err(errors.New("400")).Msgf("Bad Request, err: %s, %s", errA, errB)
	}

	return floatA, floatB, errA, errB
}

func returnA(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	a := vars["a"]
	intA, err := strconv.Atoi(a)

	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Error().Err(errors.New("400")).Msg("Bad Request")
	}

	return intA, err
}

func logAB(a float64, b float64, r *http.Request, w *http.ResponseWriter) {
	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
	log.Info().Str("a", strconv.FormatFloat(a, 'f', -1, 64)).Str("b", strconv.FormatFloat(b, 'f', -1, 64)).Msgf("IP: %s, port: %s, URL: %s", ip, port, r.URL.Path)
	(*w).WriteHeader(http.StatusOK)
}

// A Sum represents a summarize of two values.
func Sum(w http.ResponseWriter, r *http.Request) {
	a, b, errA, errB := returnAB(w, r)
	if errA == nil && errB == nil {
		logAB(a, b, r, &w)
		a += b
		fmt.Fprint(w, strconv.FormatFloat(a, 'f', -1, 64))
		return
	}

	w.WriteHeader(http.StatusBadRequest)

}

// A Diff represents a diffrence of two values.
func Diff(w http.ResponseWriter, r *http.Request) {
	a, b, errA, errB := returnAB(w, r)
	if errA == nil && errB == nil {
		logAB(a, b, r, &w)
		a -= b
		fmt.Fprint(w, strconv.FormatFloat(a, 'f', -1, 64))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

//A Div represents a division of two values.
func Div(w http.ResponseWriter, r *http.Request) {
	a, b, errA, errB := returnAB(w, r)
	if b == 0 {
		http.Error(w, "400 Bad Request", http.StatusNotFound)
		log.Error().Err(errors.New("400")).Msg("Bad Request")
		return
	} else if errA == nil || errB == nil {
		logAB(a, b, r, &w)
		a /= b
		fmt.Fprint(w, strconv.FormatFloat(a, 'f', 4, 64))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

//A Mul represents a multiplication of two values
func Mul(w http.ResponseWriter, r *http.Request) {
	a, b, errA, errB := returnAB(w, r)
	if errA == nil && errB == nil {
		logAB(a, b, r, &w)
		a *= b
		fmt.Fprint(w, strconv.FormatFloat(a, 'f', -1, 64))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

//A Fac represents a factorial of one value
func Fac(w http.ResponseWriter, r *http.Request) {
	a, err := returnA(w, r)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		ip, port, _ := net.SplitHostPort(r.RemoteAddr)
		log.Info().Str("a", strconv.Itoa(a)).Msgf("IP: %s, port: %s, URL: %s", ip, port, r.URL.Path)
		result := 1

		health.Ready.MarkAsDown()

		for i := 2; i <= a; i++ {
			result *= i
		}

		health.Ready.MarkAsUp()

		fmt.Fprint(w, strconv.Itoa(result))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
