package markdown

type ElementType int

const (
	ElementTypeParagraph ElementType = iota
	ElementTypeHeading
	ElementTypeAnchor
	ElementTypeList
)

type Element interface {
	Type() ElementType
	Markdown() string
}

type element struct {
	elementType ElementType
	markdown    string
}

func (e *element) Type() ElementType {
	return e.elementType
}

func (e *element) Markdown() string {
	return e.markdown
}

func NewElement(elementType ElementType, text string) Element {
	return &element{
		elementType: elementType,
		markdown:    text,
	}
}
