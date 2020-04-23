package messages

const (
  CodeExpansionRequestType = "code-expansion"
  CodeProcessRequestType = "code-process"
)

type CodeExpansionRequest struct {
  Reference string
}

type CodeProcessRequest struct {
  CodeToProcess string
}