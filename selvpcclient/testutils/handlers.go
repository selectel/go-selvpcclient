package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil" //nolint:staticcheck
	"net/http"
	"reflect"
	"testing"
)

// HandleReqOpts represents options for the testing utils package handlers.
type HandleReqOpts struct {
	// Mux represents HTTP Mux for a testing handler.
	Mux *http.ServeMux

	// URL represents handler's HTTP URL.
	URL string

	// RawResponse represents raw string HTTP response that needs to be returned
	// by the handler.
	RawResponse string

	// RawRequest represents raw string HTTP request that needs to be compared
	// with the actual request that will be provided by the caller.
	RawRequest string

	// Method contains HTTP method that needs to be compared against real method
	// provided by the caller.
	Method string

	// Status represents HTTP status that will be returned by the handler.
	Status int

	// CallFlag can be used to check if caller sent a request to a handler.
	CallFlag *bool
}

// HandleReqWithoutBody provides the HTTP endpoint to test requests without body.
func HandleReqWithoutBody(t *testing.T, opts *HandleReqOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		fmt.Fprint(w, opts.RawResponse)

		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		*opts.CallFlag = true
	})
}

// HandleReqWithBody provides the HTTP endpoint to test requests with body.
func HandleReqWithBody(t *testing.T, opts *HandleReqOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
		defer r.Body.Close()

		var actualRequest interface{}
		err = json.Unmarshal(b, &actualRequest)
		if err != nil {
			t.Errorf("unable to unmarshal the request body: %v", err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		fmt.Fprint(w, opts.RawResponse)

		var expectedRequest interface{}
		err = json.Unmarshal([]byte(opts.RawRequest), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw request: %v", err)
		}

		if !reflect.DeepEqual(expectedRequest, actualRequest) {
			t.Fatalf("expected %#v request, but got %#v", expectedRequest, actualRequest)
		}

		*opts.CallFlag = true
	})
}
