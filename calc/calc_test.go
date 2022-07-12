package calc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCalc(t *testing.T) {
	tests := map[string]struct {
		expectedSumStatus  int
		expectedDiffStatus int
		expectedMulStatus  int
		expectedDivStatus  int
		expectedFactStatus int
	}{
		"everything OK": {
			expectedSumStatus:  http.StatusOK,
			expectedDiffStatus: http.StatusOK,
			expectedMulStatus:  http.StatusOK,
			expectedDivStatus:  http.StatusOK,
			expectedFactStatus: http.StatusOK,
		},
		"everything BadRequest": {
			expectedSumStatus:  http.StatusBadRequest,
			expectedDiffStatus: http.StatusBadRequest,
			expectedMulStatus:  http.StatusBadRequest,
			expectedDivStatus:  http.StatusBadRequest,
			expectedFactStatus: http.StatusBadRequest,
		},
		"zero": {
			expectedSumStatus:  http.StatusOK,
			expectedDiffStatus: http.StatusOK,
			expectedMulStatus:  http.StatusOK,
			expectedDivStatus:  http.StatusBadRequest,
			expectedFactStatus: http.StatusOK,
		},
	}

	testServer := httptest.NewServer(newServer())

	defer testServer.Close()
	client := testServer.Client()

	type status struct {
		sumStatus  *http.Response
		diffStatus *http.Response
		mulStatus  *http.Response
		divStatus  *http.Response
		factStatus *http.Response
	}

	s := status{}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if name == "everything OK" {
				s.sumStatus, _ = client.Get(testServer.URL + "/sum/1/3")
				s.diffStatus, _ = client.Get(testServer.URL + "/diff/1/3")
				s.mulStatus, _ = client.Get(testServer.URL + "/mul/1/3")
				s.divStatus, _ = client.Get(testServer.URL + "/div/1/3")
				s.factStatus, _ = client.Get(testServer.URL + "/factorial/3")
			} else if name == "zero" {
				s.sumStatus, _ = client.Get(testServer.URL + "/sum/0/0")
				s.diffStatus, _ = client.Get(testServer.URL + "/diff/0/0")
				s.mulStatus, _ = client.Get(testServer.URL + "/mul/0/0")
				s.divStatus, _ = client.Get(testServer.URL + "/div/0/0")
				s.factStatus, _ = client.Get(testServer.URL + "/factorial/0")
			} else {
				s.sumStatus, _ = client.Get(testServer.URL + "/sum/abc/3")
				s.diffStatus, _ = client.Get(testServer.URL + "/diff/abc/3")
				s.mulStatus, _ = client.Get(testServer.URL + "/mul/abc/3")
				s.divStatus, _ = client.Get(testServer.URL + "/div/abc/3")
				s.factStatus, _ = client.Get(testServer.URL + "/factorial/abc")
			}

			if s.sumStatus.StatusCode != test.expectedSumStatus {
				t.Errorf("Expected: %v, got: %v", test.expectedSumStatus, s.sumStatus.StatusCode)
			}

			if s.diffStatus.StatusCode != test.expectedDiffStatus {
				t.Errorf("Expected: %v, got: %v", test.expectedDiffStatus, s.diffStatus.StatusCode)
			}

			if s.mulStatus.StatusCode != test.expectedMulStatus {
				t.Errorf("Expected: %v, got: %v", test.expectedMulStatus, s.mulStatus.StatusCode)
			}

			if s.divStatus.StatusCode != test.expectedDivStatus {
				fmt.Println(s.divStatus.Request)
				fmt.Println("---------------------------")
				fmt.Println(s.divStatus.Request.URL)
				fmt.Println(s.divStatus.StatusCode)
				//fmt.Print(s.divStatus.Request.Body)

				t.Errorf("Expected: %v, got: %v", test.expectedDivStatus, s.divStatus.StatusCode)
			}

			if s.factStatus.StatusCode != test.expectedFactStatus {
				t.Errorf("Expected: %v, got: %v", test.expectedFactStatus, s.factStatus.StatusCode)
			}
		})

	}

}

func newServer() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sum/{a}/{b}", Sum).Methods("GET")
	router.HandleFunc("/diff/{a}/{b}", Diff).Methods("GET")
	router.HandleFunc("/mul/{a}/{b}", Mul).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", Div).Methods("GET")
	router.HandleFunc("/factorial/{a}", Fac).Methods("GET")

	return router
}

/*

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
	Sum(w, req)
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
	Sum(w, req)
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
	Mul(w, req)
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
	Mul(w, req)
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
	Div(w, req)
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
	Div(w, req)
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
	Diff(w, req)
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
	Diff(w, req)
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
	Fac(w, req)
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
	Fac(w, req)
	res = w.Result()
	checkStautsCode(res, http.StatusBadRequest, t)
}

func checkStautsCode(res *http.Response, status int, t *testing.T) {
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request got %v", res.StatusCode)
	}
}
*/
