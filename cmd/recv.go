/*
Copyright © 2026 anothermeer <me@melons.cc>
*/
package cmd

import (
	"fmt"

	"github.com/anothermeer/sambal/internal/core/network"
	"github.com/spf13/cobra"
)

// recvCmd represents the recv command
var recvCmd = &cobra.Command{
	Use:   "recv",
	Short: "Listen for file transfers.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Entered receiving mode...")
		network.StartTCPSrv("3721")
	},
}

func init() {
	rootCmd.AddCommand(recvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
