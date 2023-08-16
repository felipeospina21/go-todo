/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	logging "github.com/felipeospina21/go-todo/internal"
	"github.com/felipeospina21/go-todo/todo"
	"github.com/spf13/cobra"
)

var sortFlagOpts = []string{"done", "pending"}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(DataFile)
	if err != nil {
		logging.ErrorAndQuit(err)
	}

	filter, err := cmd.Flags().GetString("filter")
	filteredItems := getFilteredVals(items, filter)

	b, err := cmd.Flags().GetBool("sorted")
	if b {
		sort.Sort(todo.ByPri(filteredItems))
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, item := range filteredItems {
		fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("sorted", "s", false, "sort by priority")
	listCmd.Flags().StringP("filter", "f", "", fmt.Sprintf("filter by %q", sortFlagOpts))
}

func getFilteredVals(items []todo.Item, filter string) []todo.Item {
	done := []todo.Item{}
	pending := []todo.Item{}
	for _, item := range items {
		if item.Done {
			done = append(done, item)
		} else {
			pending = append(pending, item)
		}
	}
	switch filter {
	case "done":
		return done
	case "pending":
		return pending
	default:
		return items

	}
}
