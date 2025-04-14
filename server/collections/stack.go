package collections

// Go doesn't have stacks, but we need a Generic stack collection for infix > postfix > prefix conversion of formulae

type Stack struct {
	items []interface{}
}

// Add element to stack
func (s *Stack) Push(data interface{}) {
	s.items = append(s.items, data)
}

func (s *Stack) Inspect() []interface{} {
	return s.items
}

// Pop element off of stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	topElement := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return topElement
}

// Check length of stack
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Check top of stack
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}
