/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	logging "github.com/felipeospina21/go-todo/internal"
	"github.com/felipeospina21/go-todo/todo"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done INDEX ...",
	Short:   "Mark items as done",
	Aliases: []string{"do"},
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		logging.ErrorAndQuit(err)
	}

	for _, arg := range args {
		i := todo.ParseArg(arg)

		if i > 0 && i < len(items) {
			items[i-1].Done = true
			items[i-1].DoneDate = time.Now()
			fmt.Printf("%v %q %v\n", items[i-1].Label(), items[i-1].Text, "marked done")

			todo.SaveItems(dataFile, items)
		} else {
			log.Println(i, "doesn't match any items")
		}

	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
