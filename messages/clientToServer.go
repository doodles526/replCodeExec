package messages

const (
  CodeExpansionRequestType = "code-expansion"
  CodeProcessRequestType = "code-process"
)

type CodeExpansionRequest struct {
  Reference string
  SessionID *string `json:"session_id,omitempty"`
}

type CodeProcessRequest struct {
  CodeToProcess string
  SessionID *string `json:"session_id,omitempty"`
}