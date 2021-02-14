package main

import (
  "context"
  "encoding/json"
  "errors"
  "log"
  "net/http"
  "strings"
  "fmt"

  "github.com/go-kit/kit/endpoint"
  httptransport "github.com/go-kit/kit/transport/http"
)

// StringService provides operations on strings.
type StringService interface {
  Uppercase(string) (string, error)
  Count(string) int
}

// stringService is a concrete implementation of StringService
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
  if s == "" {
    return "", ErrEmpty
  }
  return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
  return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// For each method, we define request and response structs
type uppercaseRequest struct {
  S string `json:"s"`
}

type uppercaseResponse struct {
  V   string `json:"v"`
  Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type countRequest struct {
  S string `json:"s"`
}

type countResponse struct {
  V int `json:"v"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
  return func(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(uppercaseRequest)
    v, err := svc.Uppercase(req.S)
    if err != nil {
      return &uppercaseResponse{v, err.Error()}, nil
    }
    return &uppercaseResponse{v, ""}, nil
  }
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
  return func(_ context.Context, request interface{}) (interface{}, error) {
    fmt.Printf("2. makeCountEndpoint\n")
    req := request.(countRequest)
    v := svc.Count(req.S)
    return &countResponse{v}, nil
  }
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
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

  // curl -X POST -H "Content-Type: application/json" -d '{"S":"abc@gmail.com"}' http://localhost:8080/uppercase
  http.Handle("/uppercase", uppercaseHandler)
  // curl -X POST -H "Content-Type: application/json" -d '{"S":"abc@gmail.com"}' http://localhost:8080/count
  http.Handle("/count", countHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request uppercaseRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    return nil, err
  }
  return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request countRequest
  fmt.Printf("1. decodeCountRequest\n")
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    return nil, err
  }
  return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  v1, ok1 := response.(*uppercaseResponse)
  v2, ok2 := response.(*countResponse)
  fmt.Printf("3. encodeResponse v1:%+v ok1=%v v2:%+v ok2=%v\n", v1, ok1, v2, ok2)
  return json.NewEncoder(w).Encode(response)
}

