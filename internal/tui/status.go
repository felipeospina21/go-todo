package tui

type status int

func (s status) getNext() status {
	if s == DONE {
		return TODO
	}
	return s + 1
}

func (s status) getPrev() status {
	if s == TODO {
		return DONE
	}
	return s - 1
}
