package tui

type Task struct {
	status      status
	title       string
	description string
	priority    int
	position    int
}

func NewTask(status status, title, description string) Task {
	return Task{status: status, title: title, description: description}
}

func (t *Task) Next() {
	if t.status == DONE {
		t.status = TODO
	} else {
		t.status++
	}
}

// implement the list.Item interface
func (t Task) FilterValue() string { return t.title }
func (t Task) Title() string       { return t.title }
func (t Task) Description() string { return t.description }
