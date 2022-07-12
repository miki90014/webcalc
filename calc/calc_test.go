package calc

import (
	"io/ioutil"
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
			expectedDivStatus:  http.StatusOK,
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

				if s.divStatus.StatusCode == http.StatusOK {
					defer s.divStatus.Body.Close()
					data, _ := ioutil.ReadAll(s.divStatus.Body)
					str := string(data)
					if str != "400 Bad Request\n" {
						t.Errorf("Expected 400 Bad Request got %v", string(data))
					}
				}
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
