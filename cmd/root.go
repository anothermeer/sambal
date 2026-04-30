/*
Copyright © 2026 anothermeer <me@melons.cc>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func runGUI() {
	println("Launching GUI...")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sambal",
	Short: "Yet another easy to use tool for sharing files and text across devices. ",
	Long: `There's a shorter version of the description so why not just look at that instead?
	tbh this section is just a bunch of random texts.. or is it?
	Yet another stupidly easy to use tool for sharing files and text across devices and platforms.
	Powered by golang, cobra, fyne, me and coffee.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		runGUI()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sambal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
