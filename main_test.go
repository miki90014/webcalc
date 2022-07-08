package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"konta.monika/webcalc/calc"
)

func TestAdd(t *testing.T) {
	req, err := http.NewRequest("GET", "/sum", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "5",
		"b": "4",
	})

	w := httptest.NewRecorder()
	calc.Sum(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "9" {
		t.Errorf("Expected 9.0000 got %v", string(data))
	}

	req, err = http.NewRequest("GET", "/sum", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "5",
		"b": "dskjf",
	})

	w = httptest.NewRecorder()
	calc.Sum(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)
}

func TestMul(t *testing.T) {
	req, err := http.NewRequest("GET", "/mul", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "1.13",
		"b": "2.80",
	})

	w := httptest.NewRecorder()
	calc.Mul(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "3.1639999999999997" {
		t.Errorf("Expected 3.1639999999999997 got %v", string(data))
	}

	req, err = http.NewRequest("GET", "/mul", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "4",
		"b": "dskjf",
	})

	w = httptest.NewRecorder()
	calc.Mul(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)

}

func TestDiv(t *testing.T) {
	req, err := http.NewRequest("GET", "/div", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "9",
		"b": "6",
	})

	w := httptest.NewRecorder()
	calc.Div(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "1.5000" {
		t.Errorf("Expected 1.5000 got %v", string(data))
	}

	req, err = http.NewRequest("GET", "/div", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "4",
		"b": "dskjf",
	})

	w = httptest.NewRecorder()
	calc.Div(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)
}

func TestDiff(t *testing.T) {
	req, err := http.NewRequest("GET", "/diff", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "5",
		"b": "4.2",
	})

	w := httptest.NewRecorder()
	calc.Diff(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "0.7999999999999998" {
		t.Errorf("Expected 0.7999999999999998 got %v", string(data))
	}

	req, err = http.NewRequest("GET", "/diff", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "4",
		"b": "dskjf",
	})

	w = httptest.NewRecorder()
	calc.Diff(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)
}

func TestFac(t *testing.T) {
	req, err := http.NewRequest("GET", "/factorial", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "5",
	})

	w := httptest.NewRecorder()
	calc.Fac(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "120" {
		t.Errorf("expected 120 got %v", string(data))
	}

	req, err = http.NewRequest("GET", "/factorial", nil)
	if err != nil {
		t.Errorf("Somthing went wrong")
	}
	req = mux.SetURLVars(req, map[string]string{
		"a": "dskjf",
	})

	w = httptest.NewRecorder()
	calc.Fac(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)
}

func checkStautsCode(res *http.Response, status int, t *testing.T) {
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request got %v", res.StatusCode)
	}
}
