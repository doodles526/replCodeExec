package execution

import (
  "os/exec"
  "strconv"
  "bufio"
  "io"

  "github.com/google/uuid"
)

type ResponseElement struct {
  // Type allows us to label the type of the payload
  Type ElementType

  // Reference allows one to point to this element for
  // potential deferred rendering
  Reference RefID

  // TODO: references by magic string aren't good. Fix that
  Representation string

  // References holds a mapping of a referenceID to a string of the yet to be parsed data
  References map[RefID]bool
}

/*
ExecutorPool
*/

type InlineExecutor interface {
  RunCodeLine(string) (*ResponseElement, error)
  GetReference(RefID) (*ResponseElement, error)
}

type pythonExecutor struct {
  cmd *exec.Cmd
  stdin io.WriteCloser
  stdout io.ReadCloser
  stderr io.ReadCloser

  stdoutB bufio.Reader
  stderrB bufio.Reader

  refState map[RefID]ResponseElement
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
    stdoutB: bufio.NewReader(stdout),
    stderrB: bufio.NewReader(stderr),
    refState: map[RefID]ResponseElement{},
  }
}

func (p *pythonExecutor) GetReference(ref RefID) (*ResponseElement, error) {
  resp, ok := p[ref]
  if !ok {
    return nil, fmt.ErrorF("Could not find reference %v", ref)
  }
  return resp, nil
}
// TODO: For python environment, implement python stub script that allows receiving lines over STDIN and responding over STDOUT
func(p *pythonExecutor) RunCodeLine(line string) (*ResponseElement, error) {
  if _, err := io.WriteString(p.stdin, line); err != nil {
    return nil, err
  }

  interpResp, err := p.stdoutB.ReadLine()
  if err != nil {
    return nil, err
  }
  return p.parse(interpResp)
} 

// parse should parse returned data. It should return a responseelement
func (p *pythonExecutor) parse(code string) (*ResponseElement, error) {
  switch p.topLevelDataType(code) {
  case StringElementType:
    p.parseString(code)
  case ObjectElementType:
    p.parseObject(code)
  case NumberElementType:
    p.parseNumber(code)
  case ArrayElementType:
    p.parseArray(code)
  default:
    return nil, fmt.Errorf("Unknown data type")
  }
  return nil, fmt.Errorf("Should have found recognized type")
}

func (p *pythonExecutor) parseNumber(code string) (*ResponseElement, error) {
  // Just validation logic
  val, err := strconv.Atoi(code)
  if err != nil {
    return nil, err
  }
  return &ResponseElement {
    Type: NumberElementType,    
    Reference: uuid.New(),
    Representation: code,
  }, nil
}

func (p *pythonExecutor) parseString(code string) (*ResponseElement, error) {
  return &ResponseElement {
    Type: StringElementType,
    Reference: uuid.New(),
    Representation: code,
  }, nil
}

func (p *pythonExecutor) parseArray(code string) (*ResponseElement, error) {
  // TODO: Don't assume proper formatting
  contents = strings.TrimSpace(code)
  contents = strings.[1:len(contents)-2]
  elements = strings.Split(contents, ',')

  refs := map[RefID]bool{}

  finalRep := "["
  representationElements := []string{}
  for _, elem := range elements {
    if p.shouldRefer(elem) {
      rElem, err := p.parse(elem)
      if err != nil {
        return nil, err
      }
      // Cache the reference for later expansion
      p[rElem.Reference] = rElem
      // have replacement refs in place
      refs[rElem.Reference] = true
      representationElements = append(representationElements, string(rElem.Reference))

    } else {
      representationElements = append(representationElements, elem) 
    }
  }
  finalRep += strings.Join(representationElements, ", ")
  finalRep += "]"
  
  return &ResponseElement{
    Type: ArrayElementType,
    Reference: UUID.New(),
    Representation: finalRep,
    References: refs,
  }, nil
}

// TODO extract code for representation for code dupe with parseArray
func (p *pythonExecutor) parseObject(code string) (*ResponseElement, error) {
  contents = strings.TrimSpace(code)
  contents = strings.[1:len(contents)-2]
  elements = strings.Split(contents, ',')

  refs := map[RefID]bool{}

  finalRep := "{"
  representationElements := []string{}

  for _, elem := range elements {
    keyValPair := strings.Split(elem, ':')    
    key := keyValPair[0]
    key = strings.TrimSpace(key)
    // assume key is primitive type

    val := keyValPair[1]
    val = strings.TrimSpace(val)

    if p.shouldRefer(val) {
      rElem, err := p.parse(val)
      if err != nil {
        return nil, err
      }
      // Cache the reference for later expansion
      p[rElem.Reference] = rElem
      // have replacement refs in place
      refs[rElem.Reference] = true
      kvPair := key + ":" + string(rElem.Reference)
      representationElements = append(representationElements, kvPair)

    } else {
      kvPair := key + ": " + val
      representationElements = append(representationElements, elem) 
    }
  }
  return &ResponseElement{
    Type: ObjectElementType,
    Reference: UUID.New(),
    Representation: finalRep,
    References: refs,
  }, nil
}

func (p *pythonExecutor) shouldRefer(code string) bool{
  switch p.topLevelDataType(code) {
  case StringElementType:
    return false
  case ObjectElementType:
    return true
  case NumberElementType:
    return false
  case ArrayElementType:
    return true
  // Don't refer to something if we don't know what it is
  return false
}

func (p *pythonExecutor) topLevelDataType(code string) ElementType {
  firstChar := code[0]
  switch firstChar {
  case '\'':
    return StringElementType
  case '{':
    return ObjectElementType
  case '[':
    return ArrayElementType
  default: // for now assume no char prefix means it is a number
    return NumberElementType
  }

}