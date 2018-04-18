package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

// HandleReqWithoutBody provides the HTTP endpoint to test requests without body.
func HandleReqWithoutBody(mux *http.ServeMux, url, rawResponse, method string, httpStatus int, callFlag *bool, t *testing.T) {
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		fmt.Fprintf(w, rawResponse)

		if r.Method != method {
			t.Fatalf("expected %s method but got %s", method, r.Method)
		}

		*callFlag = true
	})
}

// HandleReqWithBody provides the HTTP endpoint to test requests with body.
func HandleReqWithBody(mux *http.ServeMux, url, rawResponse, rawRequest, method string, httpStatus int, callFlag *bool, t *testing.T) {
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		fmt.Fprintf(w, rawResponse)

		if r.Method != method {
			t.Fatalf("expected %s method but got %s", method, r.Method)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}

		var actualRequest interface{}
		err = json.Unmarshal(b, &actualRequest)
		if err != nil {
			t.Errorf("unable to unmarshal the request body: %v", err)
		}

		var expectedRequest interface{}
		err = json.Unmarshal([]byte(rawRequest), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw request: %v", err)
		}

		if !reflect.DeepEqual(expectedRequest, actualRequest) {
			t.Fatalf("expected %#v request, but got %#v", expectedRequest, actualRequest)
		}

		*callFlag = true
	})
}
