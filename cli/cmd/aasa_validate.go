package cmd

import (
	"fmt"

	"github.com/BlinqOSS/yurl/yurllib"

	"github.com/spf13/cobra"
)

var filePath string

// validateAASACmd represents the validate command for Apple App Site Association
var validateAASACmd = &cobra.Command{
	Use:   "validate [<URL>]",
	Short: "Validate your link or local file against Apple's requirements",
	Run: func(cmd *cobra.Command, args []string) {
		if filePath != "" {
			output, err := yurllib.CheckAASAFile(filePath)

			if err != nil {
				fmt.Println("Got error: %w", err)
			}

			for _, item := range output {
				fmt.Print(item)
			}

			return
		} else {
			output := yurllib.CheckAASADomain(args[0], "", "", true)

			for _, item := range output {
				fmt.Print(item)
			}
		}
	},
}

func init() {
	validateAASACmd.Flags().StringVarP(&filePath, "file", "f", "", "path to local file, will ignore the URL")
	aasaCmd.AddCommand(validateAASACmd)
}
