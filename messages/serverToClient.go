package messages

const (
	CodeResponseGraphicalType = "resp-graphical"
	CodeResponseTerminalType  = "resp-terminal"
	CodeResponseErrorType     = "resp-error"
)

// GraphicalCodeResponse should result in a ref to an image
// To render on screen
type GraphicalCodeResponse struct {
	Reference string  `json:"ref"`
	Location  string  `json:"location"`
	SessionID *string `json:"session_id,omitempty"`
}

type TerminalCodeResponse struct {
	Text                        string            `json:"text"`
	ReservedExpansionReferences map[string]string `json:"reserved_refs"`
	SessionID                   *string           `json:"session_id,omitempty"`
}

/*
{
  text: "{'a': @@MAGIC_STRING_1}"
  reserved_refs: {'@@MAGIC_STRING_1': '...'}
}
*/
type CodeResponseError struct {
	Error string `json:"error"`
}
