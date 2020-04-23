package messages

const (
  CodeResponseGraphicalType = "resp-graphical" 
  CodeResponseTerminalType = "resp-terminal" 
)

// GraphicalCodeResponse should result in a ref to an image
// To render on screen
type GraphicalCodeResponse struct {
  Reference string `json:"ref"`
  Location string `json:"location"`
}

type TerminalCodeResponse struct {
  Text string `json:"text"`
  ReservedExpansionReferences map[string]string `json:"reserved_refs"`
}

