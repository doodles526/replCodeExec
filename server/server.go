package server

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/"
)

type Args struct {
  ListenAddr *net.TCPAddr
  DefaultClientConfig *ClientConfig
}

func ServeCodeExecution(args *Args) error {
  // TODO: Pass a code execution client in
  r := mux.NewRouter()

  // TODO: Do I want this to be a single-wire API?
  // TODO: Wire eventual execution client into factory
  r.HandleFunc("/", executionHandlerFactory()).Methods("POST")
}

func executionHandlerFactory() func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
  }
}