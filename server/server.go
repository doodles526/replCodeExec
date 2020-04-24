package server

import (
  "net/http"
  "io/ioutil"

  "github.com/gorilla/mux"
  "github.com/doodles526/replCodeExec/messages"
  "github.com/doodles526/replCodeExec/execution"
)

type Args struct {
  ListenAddr string
  Executor execution.Executor
}

func ServeCodeExecution(args *Args) error {
  // TODO: Pass a code execution client in
  r := mux.NewRouter()


  // TODO: Do I want this to be a single-wire API?
  // TODO: Wire eventual execution client into factory
  r.HandleFunc("/", executionHandlerFactory(args.Executor)).Methods("POST")
  s := &http.Server {
    Addr: args.ListenAddr,
    Handler: r,
  }
  return s.ListenAndServe()
}

func executionHandlerFactory(execPool execution.Executor) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      errResp := messages.CodeResponseError{
        Error: err.Error(),//TODO: Obscure error for security
      }
      sendError(errResp, 400, w)
      return
    }

    var codeReq messages.Message
    msgType, msg, err := messages.UnmarshalMsg(body)

    switch msgType {
    case messages.CodeProcessRequestType:
      handleCodeProcessRequest(execPool, w, r, msg)
    case messages.CodeExpansionRequestType:
      handleCodeExpansionRequest(execPool, w, r, msg)
    default:
      errResp := message.CodeResponseError {
        Error: "Unknown request type"
      }
      sendError(errResp, 400, w)
    }
  }
}

func sendError(msg messages.CodeResponseError, code int, w http.ResponseWriter) {
      w.WriteHeader(code)
      errRData, err := messages.MarshalMessage(errResp)
      if err != nil {
        w.WriteHeader(500)
        fmt.Fprintf(w, "Error encountered, but could not render error message.")
        return
      }
      w.Write(errRData)
}

func handleCodeProcessRequest(exec execution.Executor, w http.ResponseWriter, r *http.Request, msg messages.CodeProcessRequest) {
  // TODO implement
  execResp, err := exec.RunCodeLine(msg.CodeToProcess)
  if err != nil {
    // TODO
  }
  clientRef := map[string]string {}
  for ref, _ := range execResp.References {
    // what to replace references with in display
    clientRef[string(ref)] = "..."
  }
  clientResp := &TerminalCodeResponse{
    Text: execResp.Representation,
    ReservedExpansionReferences clientRef,
    // TODO: Pass in session identifier
    SessionID: nil,
  }
  msg, err := messages.Marshal(clientResp)
  if err != nil {
    // TODO
  }
  w.Write(msg)
}

func handleCodeExpansionRequest(exec execution.Executor, w http.ResponseWriter, r *http.Request, msg messages.CodeExpansionRequest) {
  ref, err := exec.GetReference([]byte(msg.Reference))
  if err != nil {
    // TODO
  }
  clientRef := map[string]string {}
  for ref, _ := range ref.References {
    // what to replace references with in display
    clientRef[string(ref)] = "..."
  }

  clientResp := &TerminalCodeResponse{
    Text: execResp.Representation,
    ReservedExpansionReferences clientRef,
    // TODO: Pass in session identifier
    SessionID: nil,
  }

  msg, err := messages.Marshal(clientResp)
  if err != nil {
    // TODO
  }
  w.Write(msg)

  // TODO implement
}