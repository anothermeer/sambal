/*
Copyright © 2026 anothermeer <me@melons.cc>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/anothermeer/sambal/internal/core/discovery"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List nearby devices.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Scanning devices...")
		devices, err := discovery.Browse(5 * time.Second)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(devices) == 0 {
			fmt.Println("No Sambal devices found.")
			return
		}

		for _, d := range devices {
			fmt.Printf(
				"%-18s %-15s %-8s %s\n",
				d.Name,
				d.Addr,
				d.Version,
				d.ID,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
