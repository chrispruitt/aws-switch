package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	_ "github.com/chrispruitt/aws-switch/state"
	"github.com/spf13/cobra"
)

var (
	tagsInput []string
)

func init() {
	HaltCmd.PersistentFlags().StringArrayVar(&tagsInput, "tag", nil, "tags to identify resources to halt. \"key=value\"")
}

var HaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "halt an aws service",
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

		lib.Halt(tags)
	},
}
