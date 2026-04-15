package layout

type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	l := len(s.values)
	if l == 0 {
		panic("nothing on the stack left to pop")
	}
	value := s.values[l-1]
	s.values = s.values[:l-1]
	return value
}

func (s *stack) pop2() (int, int) {
	l := len(s.values)
	if l <= 1 {
		panic("not enough on the stack left to pop")
	}
	v1, v2 := s.values[l-1], s.values[l-2]
	s.values = s.values[:l-2]
	return v1, v2
}

func (s *stack) sum() {
	v1, v2 := s.pop2()
	s.push(v2 + v1)
}

func (s *stack) diff() {
	v1, v2 := s.pop2()
	s.push(v2 - v1)
}

func (s *stack) prod() {
	v1, v2 := s.pop2()
	s.push(v2 * v1)
}

func (s *stack) div() {
	v1, v2 := s.pop2()
	s.push(v2 / v1)
}

func (s *stack) mod() {
	v1, v2 := s.pop2()
	s.push(v2 % v1)
}

func (s *stack) eq() {
	v1, v2 := s.pop2()
	isEqual := 0
	if v2 == v1 {
		isEqual = 1
	}
	s.push(isEqual)
}

func (s *stack) gt() {
	v1, v2 := s.pop2()
	isGreaterThan := 0
	if v2 > v1 {
		isGreaterThan = 1
	}
	s.push(isGreaterThan)
}

func (s *stack) lt() {
	v1, v2 := s.pop2()
	isGreaterThan := 0
	if v2 < v1 {
		isGreaterThan = 1
	}
	s.push(isGreaterThan)
}

func (s *stack) dup() {
	l := len(s.values)
	if l == 0 {
		panic("nothing on the stack left to dup")
	}
	s.values = append(s.values, s.values[l-1])
}

func (s *stack) rm() {
	_ = s.pop()
}
