package gostack

type Stack struct {
	elements []ASTElement
}

func NewStack(tokenizer *Tokenizer) *Stack {
	p := new(Stack)
	return p
}

func (s *Stack) Push(e ASTElement) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() ASTElement {
	len := len(s.elements)
	e := s.elements[len-1]
	s.elements = s.elements[:len-1]
	return e
}

func (s *Stack) Top() ASTElement {
	return s.elements[len(s.elements)-1]
}
