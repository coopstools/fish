package layout

type stackContainer struct {
	base *stack
}

type stack struct {
	parent   *stack
	register *int

	values []int
}

func (s *stackContainer) push(v int) {
	s.base.values = append(s.base.values, v)
}

func (s *stackContainer) pop() int {
	l := len(s.base.values)
	if l == 0 {
		panic("nothing on the stack left to pop")
	}
	value := s.base.values[l-1]
	s.base.values = s.base.values[:l-1]
	return value
}

func (s *stackContainer) pop2() (int, int) {
	l := len(s.base.values)
	if l <= 1 {
		panic("not enough on the stack left to pop")
	}
	v1, v2 := s.base.values[l-1], s.base.values[l-2]
	s.base.values = s.base.values[:l-2]
	return v1, v2
}

func (s *stackContainer) pop3() (int, int, int) {
	l := len(s.base.values)
	if l <= 2 {
		panic("not enough on the stack left to pop")
	}
	v1, v2, v3 := s.base.values[l-1], s.base.values[l-2], s.base.values[l-3]
	s.base.values = s.base.values[:l-3]
	return v1, v2, v3
}

func (s *stackContainer) sum() {
	v1, v2 := s.pop2()
	s.push(v2 + v1)
}

func (s *stackContainer) diff() {
	v1, v2 := s.pop2()
	s.push(v2 - v1)
}

func (s *stackContainer) prod() {
	v1, v2 := s.pop2()
	s.push(v2 * v1)
}

func (s *stackContainer) div() {
	v1, v2 := s.pop2()
	s.push(v2 / v1)
}

func (s *stackContainer) mod() {
	v1, v2 := s.pop2()
	s.push(v2 % v1)
}

func (s *stackContainer) eq() {
	v1, v2 := s.pop2()
	isEqual := 0
	if v2 == v1 {
		isEqual = 1
	}
	s.push(isEqual)
}

func (s *stackContainer) gt() {
	v1, v2 := s.pop2()
	isGreaterThan := 0
	if v2 > v1 {
		isGreaterThan = 1
	}
	s.push(isGreaterThan)
}

func (s *stackContainer) lt() {
	v1, v2 := s.pop2()
	isGreaterThan := 0
	if v2 < v1 {
		isGreaterThan = 1
	}
	s.push(isGreaterThan)
}

func (s *stackContainer) dup() {
	l := len(s.base.values)
	if l == 0 {
		panic("nothing on the stack left to dup")
	}
	s.base.values = append(s.base.values, s.base.values[l-1])
}

func (s *stackContainer) rm() {
	_ = s.pop()
}

func (s *stackContainer) swap() {
	v1, v2 := s.pop2()
	s.push(v1)
	s.push(v2)
}

func (s *stackContainer) shift3R() {
	v1, v2, v3 := s.pop3()
	s.push(v1)
	s.push(v3)
	s.push(v2)
}

func (s *stackContainer) shiftR() {
	if len(s.base.values) <= 1 {
		return
	}
	last := s.pop()
	s.base.values = append([]int{last}, s.base.values...)
}

func (s *stackContainer) shiftL() {
	if len(s.base.values) <= 1 {
		return
	}
	first := s.base.values[0]
	s.base.values = append(s.base.values[1:], first)
}

func (s *stackContainer) reverse() {
	var newStack []int
	for _, v := range s.base.values {
		newStack = append(newStack, v)
	}
	s.base.values = newStack
}

func (s *stackContainer) length() {
	s.push(len(s.base.values))
}

func (s *stackContainer) substack() {
	numToPull := s.pop()
	parentLength := len(s.base.values)
	subValues := s.base.values[parentLength-numToPull : parentLength]
	s.base.values = s.base.values[:parentLength-numToPull]
	subStack := &stack{
		parent: s.base,
		values: subValues,
	}
	s.base = subStack
}

func (s *stackContainer) restack() {
	parent := s.base.parent
	parent.values = append(parent.values, s.base.values...)
	s.base = parent
}

func (s *stackContainer) reg() {
	if s.base.register == nil {
		v := s.pop()
		s.base.register = &v
		return
	}
	s.push(*s.base.register)
	s.base.register = nil
}
