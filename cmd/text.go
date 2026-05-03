/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/anothermeer/sambal/internal/core/actions"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		actions.SendText("localhost:3721", args[0])
	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
