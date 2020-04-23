package execution



type InlineExecutor interface {
  RunCodeLine(string) (CodeResponse, error)

}

type pythonExecutor struct {

}

func(p *pythonExecutor) RunCodeLine(line string) (CodeResponse, error) {

} 

type mockExecutor struct {

}