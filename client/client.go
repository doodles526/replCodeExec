package client

import (
  "net/http"
  "net"
)

type CodeExecutionClient struct {
  c *http.Client
}

type CodeClientArgs struct {
  Remote *net.TCPAddr
  ExecConfig *ExecutionConfig 
}