/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	logging "github.com/felipeospina21/go-todo/internal"
	"github.com/felipeospina21/go-todo/todo"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del INDEX ...",
	Short: "Delete todos",
	Run:   delRun,
}

func delRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		logging.ErrorAndQuit(err)
	}

	idx := []int{}
	newItems := []todo.Item{}

	for _, arg := range args {
		i := todo.ParseArg(arg)
		idx = append(idx, i-1)
	}

	for i, item := range items {

		b := slices.Contains(idx, i)

		if !b {
			newItems = append(newItems, item)
		} else {
			fmt.Printf("%v %q %v\n", items[i].Label(), items[i].Text, "removed")
		}
	}
	todo.SaveItems(dataFile, newItems)
}

func init() {
	rootCmd.AddCommand(delCmd)
}
