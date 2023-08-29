package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/mitchellh/go-homedir"
)

// Provides the mock data to fill the kanban board

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(TODO),
		newColumn(IN_PROGRESS),
		newColumn(DONE),
	}
	home, _ := homedir.Dir()
	dataFile := home + string(os.PathSeparator) + "todos.json"
	items, _ := ReadItems(dataFile)

	// Init To Do
	b.cols[TODO].list.Title = "To Do"
	b.cols[TODO].list.SetItems(items)
	// b.cols[TODO].list.SetItems([]list.Item{
	// 	Task{status: TODO, title: "buy milk", description: "strawberry milk"},
	// 	Task{status: TODO, title: "eat sushi", description: "negitoro roll, miso soup, rice"},
	// 	Task{status: TODO, title: "fold laundry", description: "or wear wrinkly t-shirts"},
	// })
	// Init in progress
	b.cols[IN_PROGRESS].list.Title = "In Progress"
	b.cols[IN_PROGRESS].list.SetItems([]list.Item{
		Task{status: IN_PROGRESS, title: "write code", description: "don't worry, it's Go"},
	})
	// Init done
	b.cols[DONE].list.Title = "Done"
	b.cols[DONE].list.SetItems([]list.Item{
		Task{status: DONE, title: "stay cool", description: "as a cucumber"},
	})
}
