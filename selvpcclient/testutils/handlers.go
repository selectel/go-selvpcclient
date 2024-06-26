package testutils

import (
	"encoding/json"
	"html/template"
	"io"
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

		tpl, err := template.New("template").Parse(opts.RawResponse)
		if err != nil {
			t.Fatalf("unable to parse template: %v", err)
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			t.Fatalf("unable to write response: %v", err)
		}

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

		_, contentTypeExists := r.Header["Content-Type"]
		if !contentTypeExists {
			t.Fatalf("request doesn't contain content-type in headers")
		}

		contentTypeFound := false

		for _, contentTypeValue := range r.Header["Content-Type"] {
			if contentTypeValue == "application/json" {
				contentTypeFound = true

				break
			}
		}

		if !contentTypeFound {
			t.Fatalf("content-type is not equal to application/json")
		}

		b, err := io.ReadAll(r.Body)
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

		tpl, err := template.New("template").Parse(opts.RawResponse)
		if err != nil {
			t.Fatalf("unable to parse template: %v", err)
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			t.Fatalf("unable to write response: %v", err)
		}

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
