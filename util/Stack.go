package util

type Stack []string

func (s *Stack) Push(input string) {
	*s = append(*s, input)
}

func (s *Stack) Pop() *string {
	if s.IsEmpty() {
		return nil
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &result
}

func (s *Stack) Peak() *string {
	if s.IsEmpty() {
		return nil
	}
	return &(*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
