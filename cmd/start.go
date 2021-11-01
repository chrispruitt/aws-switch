package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	"github.com/spf13/cobra"
)

func init() {
	StartCmd.PersistentFlags().StringArrayVar(&tagsInput, "tag", nil, "tags to identify resources to start. \"key=value\"")
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a halted aws service",
	Run: func(cmd *cobra.Command, args []string) {

		if len(tagsInput) == 0 {
			fmt.Println("--tags required")
			os.Exit(1)
		}

		tags, err := parseTagsFlag(tagsInput)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lib.Start(tags)
	},
}
