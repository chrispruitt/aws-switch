package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	"github.com/spf13/cobra"
)

var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Creates an s3 bucket for the aws-switch state to reside.",
	Run: func(cmd *cobra.Command, args []string) {

		bucket, err := lib.Configure()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Created and configured bucket \"s3://%s\"\n", bucket)
		}
	},
}
