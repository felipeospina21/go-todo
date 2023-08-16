/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/felipeospina21/go-todo/internal/tui"
	"github.com/spf13/cobra"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Execute tui",
	Run: func(cmd *cobra.Command, args []string) {
		tui.InitBoard()
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
