package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

/*
	==========================
	BUSINESS LOGIC
	==========================
*/
var ErrEmptyString = errors.New("empty string")

// StringService provides operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// stringService is a concrete implementation of StringService
type stringService struct{}

/* the next two methods satisfy the internal StringService interface */
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

/*
	==========================
	REQUESTS & RESPONSES
	==========================
*/

// for each method, define a request & response struct
// NOTE: 'S' = "string"; 'V' = "value"
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err"` // 'errors' types don't define JSON marshaling, so use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

/*
	==========================
	ENDPOINTS
	==========================
*/

// Endpoints are a primary abstraction, representing a single RPC (method in the above-defined service interface)
// defined as:
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{V: v, Err: err.Error()}, nil
		}

		return uppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)

		return countResponse{V: v}, nil
	}
}

/*
	==========================
	TRANSPORTS
	==========================
*/

// transports expose the service to the network
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8882", nil))
}
