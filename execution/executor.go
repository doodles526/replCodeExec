package execution

import (
  "os/exec"
  "bufio"
  "io"
)

type ResponseElement struct {
  // Type allows us to label the type of the payload
  Type ElementType

  // Reference allows one to point to this element for
  // potential deferred rendering
  Reference RefType

  // TODO: Brainstorm not using an interface here
  Payload interface{}
}

func (re *ResponseElement) RenderToString(args *ResponseRenderArgs) (string, error) {
  switch re.Type {
  case ObjectElementType:
    return re.renderObject(re.Payload)
  case NumberElementType:
  case StringElementType:
  case ArrayElementType:
  case KeyValueElementType:
  default:
    return "", fmt.Errorf("Unknown element type %s", re.Type)
  }

}



type InlineExecutor interface {
  RunCodeLine(string) (ResponseElement, error)
  GetReference(RefType) (ResponseElement, error)
}

type pythonExecutor struct {
  cmd *exec.Cmd
  stdin io.WriteCloser
  stdout io.ReadCloser
  stderr io.ReadCloser

  stdoutB bufio.Reader
  stderrB bufio.REader
}

func NewPythonExecutor() (*pythonExecutor, error){
  c := exec.Command("/usr/bin/python")
  stdin, err := c.StdinPipe()
  if err != nil {
    return nil, err
  }
  stdout, err := c.StdoutPipe()
  if err != nil {
    return nil, err
  }
  stderr, err := c.StderrPipe()
  if err != nil{
    return nil, err
  }
  go c.Run()
  return &pythonExecutor{
    cmd: c,
    stdin: stdin,
    stdout: stdout,
    stderr: stderr,
    stdoutB: bufio.NewReader(stdout)
    stderrB: bufio.NewReader(stderr)
  }
}

// TODO: For python environment, implement python stub script that allows receiving lines over STDIN and responding over STDOUT
func(p *pythonExecutor) RunCodeLine(line string) (*ResponseElement, error) {
  if _, err := io.WriteString(p.stdin, line); err != nil {
    return nil, err
  }

  // TODO: Likely won't work as we need to encounter EOF - try out
  interpResp, err := ioutil.ReadAll()
  if err != nil {
    return nil, err
  }
  return p.parse(interpResp)
} 

