package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	logging "github.com/felipeospina21/go-todo/internal"
)

type Item struct {
	Text     string
	Priority int
	Done     bool
	DoneDate time.Time
	position int
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	items := []Item{}

	if err != nil {
		if f, c_err := os.Create(filename); c_err != nil {
			return items, c_err
		} else {
			defer f.Close()
		}
	}

	if err := json.Unmarshal(b, &items); err != nil {
		return items, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func DeleteItems(items []Item) {
}

func ParseArg(arg string) int {
	i, err := strconv.Atoi(arg)
	if err != nil {
		logging.ErrorAndQuitF(err, fmt.Sprintf("%s is not a valid index\n", arg))
	}

	return i
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}

	return ""
}

type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}

	return s[i].Priority < s[j].Priority
}
