package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	"github.com/spf13/cobra"
)

func init() {
	ResumeCmd.PersistentFlags().StringArrayVar(&tagsInput, "tag", nil, "tags to identify resources to resume. \"key=value\"")
}

var ResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume a halted aws service",
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

		lib.Resume(tags)
	},
}
