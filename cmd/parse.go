/*
Copyright © 2024 Mike Borozdin mikebz@
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var directory string

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse vCard files in a directory",
	Long: `
Use the parse command to take a directory of vCard files and output a CSV
based on the contents of the vCard files.  All the files that do not have
an email address will be ignored.  When the vCard has just one email address
that emails address will be used.  When the vCard has multiple email addresses
the first email address will be used unless the preferred type is specified. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("parse called, directory is %s\n", directory)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Help message for toggle")
}
