/*
Copyright © 2024 Mike Borozdin mikebz@
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mikebz/vcardlist/internal/parse"
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

		// validate that the directory parameter indeed points to a valid
		// directory.
		fileInfo, err := os.Stat(directory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Directory %s does not exist", directory)
			return
		}

		if !fileInfo.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is not a directory", directory)
			return
		}

		// parse the directory and output the CSV
		nameEmailArray, err := parse.Parse(directory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing directory %s", directory)
			return
		}

		fmt.Println("Name,Email")

		for _, nameEmail := range *nameEmailArray {
			fmt.Printf("%s,%s\n", nameEmail.Name, nameEmail.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Help message for toggle")
}
