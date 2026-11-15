package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	output := IntList{}
	for _, v := range s {
		if fn(v) {
			output = append(output, v)
		}
	}
	return output
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	output := IntList{}
	for _, v := range s {
		output = append(output, fn(v))
	}
	return output
}

func (s IntList) Reverse() IntList {
	output := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		output = append(output, s[i])
	}
	return output
}

func (s IntList) Append(lst IntList) IntList {
	for _, v := range lst {
		s = append(s, v)
	}
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = s.Append(lst)
	}
	return s
}
