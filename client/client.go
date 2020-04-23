package client

import (
  "net/http"
  "bytes"
  "io/ioutil"

  "github.com/doodles526/replCodeExec/messages"
)

type CodeExecutionClient struct {
  c *http.Client
  sAddr string
  sess *string
}

type CodeClientArgs struct {
  Remote string
  ExecConfig *ExecutionConfig 
}

func NewClient(args *CodeClientArgs) (*CodeExecutionClient, error) {
  c := &http.Client{}
  return &CodeExecutionClient{
    c: c,
    sAddr: args.Remote,
  }, nil
}

type RunCodeArgs struct {
  CodeLine string
}
// TODO: Should I process the response before return to something prettier?
func (c *CodeExecutionClient) RunCode(args *RunCodeArgs)error {
  msg := CodeProcessRequest {
    CodeToProcess: args.CodeLine,
    SessionID: c.sess,
  }
  respMsgData, err := c.ExecuteRequest(msg)
  if err != nil {
    return err
  }
  msgType, msg, err := messages.MarshalMessage(respMsgData)
  if err != nil {
    return nil, err
  }
  switch msgType {
  case messages.CodeResponseGraphicalType:
    return c.handleGraphicalResponse(msg)
  case messages.CodeResponseTerminalType:
    return c.handleTerminalResponse(msg)
  default:
    return fmt.Errorf("Unknown response type from server")
  }

  return fmt.Errorf("Should have resolved before here")
}

type ExpandReferenceArgs struct {
  Reference string
}
func (c *CodeExecutionClient)ExpandReference(args *ExpandReferenceArgs) error {
  req := messages.CodeExpansionRequest {
    Reference: args.Reference,
    SessionID: c.sess,
  }
  respMsgData, err := c.ExecuteRequest(msg)
  if err != nil {
    return err
  }
  msgType, msg, err := messages.MarshalMessage(respMsgData)
  if err != nil {
    return nil, err
  }
  switch msgType {
  case messages.CodeResponseTerminalType:
    return c.handleTerminalResponse(msg)
  default:
  // doesn't make sense for a graphical response to a expansion ref
    return fmt.Errorf("Unexpected message type %s", msgType)
  return fmt.Errorf("Should have resolved before here")
}

func(c *CodeExecutionClient) ExecuteRequest(msg interface{}) ([]byte, error){

  msgData, err := messages.MarshalMessage(msg)
  if err != nil {
    return nil, err
  }

  req, err : http.NewRequest("POST", sAddr, bytes.NewBuffer(msgData))
  if err != nil {
    return nil, err
  }
  resp, err := c.c.Do(req)
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}

func(c *CodeExecutionClient)handleGraphicalResponse(msg *messages.GraphicalCodeResponse) error {
  fmt.Println(msg)
}
func(c *CodeExecutionClient)handleTerminalResponse(msg *messages.GraphicalCodeResponse) error {
  fmt.Println(msg)
}
  case messages.CodeResponseTerminalType:
    return c.handleTerminalResponse(msg)