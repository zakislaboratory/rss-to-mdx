package markdown

type ElementType int

const (
	ElementTypeParagraph ElementType = iota
	ElementTypeHeading
	ElementTypeAnchor
	ElementTypeList
	ElementTypeListItem
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

func (e *element) Children() []Element {
	return nil
}

func NewElement(elementType ElementType, text string) Element {
	return &element{
		elementType: elementType,
		markdown:    text,
	}
}
