// Package cmd is the entry point for the application.
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "url-extractor",
	Short: "A simple ETL tool",
	Long:  `A simple ETL tool that allows you to run consume URLs`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
