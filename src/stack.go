package gostack

type ElementType int

type Stack struct {
	Elements []ElementType
}

func NewStack() *Stack {
	p := new(Stack)
	p.Elements = make([]ElementType, 0)
	return p
}

func (s *Stack) Push(e ElementType) {
	s.Elements = append([]ElementType{e}, s.Elements...)
}

func (s *Stack) Pop() ElementType {
	if len(s.Elements) == 0 {
		return 0
	}
	e := s.Elements[0]
	s.Elements = s.Elements[1:]
	return e
}

func (s *Stack) Top() ElementType {
	if len(s.Elements) == 0 {
		return 0
	}
	return s.Elements[0]
}
