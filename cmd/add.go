/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"

	logging "github.com/felipeospina21/go-todo/internal"
	"github.com/felipeospina21/go-todo/todo"
	"github.com/spf13/cobra"
)

var (
	priority  int
	pFlagVals = []int{1, 2, 3}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add TASK ...",
	Short: "Add new tasks",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		logging.ErrorAndQuit(fmt.Errorf("Please provide a todo"))
	}

	v, _ := cmd.Flags().GetInt("priority")
	checkPriorityFlagVal(v)

	items, err := todo.ReadItems(dataFile)

	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
		fmt.Printf("%q added\n", item.Text)
	}

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		logging.ErrorAndQuit(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, fmt.Sprintf("Priority:%v", pFlagVals))
}

func checkPriorityFlagVal(flagVal int) {
	if !slices.Contains(pFlagVals, flagVal) {
		fmt.Printf("invalid -p flag value %v, please choose one value from %v", flagVal, pFlagVals)
		os.Exit(1)
	}
}
