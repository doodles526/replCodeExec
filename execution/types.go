package execution

type RefID []byte
type ElementType int
const (
  ObjectElementType = iota
  NumberElementType
  StringElementType
  ArrayElementType
  // KeyValueElement allows us 
  // to have nested elements for our Object type
  KeyValueElementType
)

// ObjectElement contains a set of references to Key/Val
// Pairs. This allows us to keep constant time lookup
// while emulating dynamic key types
// TODO: Can embedded interfaces actually work here?
type ObjectElement map[ResponseElement]ResponseElement

type NumberElement int
type StringElement int
type ArrayElement []ResponseElement
type KeyValueElement struct {
  Key ResponseElement
  Val ResponseElement
}