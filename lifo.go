package lifo

// Element represents an element of the stack
type Element struct {
	name  string
	value string
}

// Stack holds all information about the stack
type Stack struct {
	Elements []Element
}

// Pop removes last element from the stack
func (st *Stack) Pop() Element {
	// TODO: removes last element from the array
	return Element{}
}

// Push serves to put an element in the stack
func (st *Stack) Push(el Element) error {
	// TODO: pushes an element to the stack
	return nil
}
