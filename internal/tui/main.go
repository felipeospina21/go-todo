package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	MARGIN = 4
	APPEND = -1
)

const (
	TODO status = iota
	IN_PROGRESS
	DONE
)

var board *Board

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
}

func InitBoard() {
	board = NewBoard()
	board.initLists()
	p := tea.NewProgram(board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
