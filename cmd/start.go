package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	"github.com/spf13/cobra"
)

func init() {
	ResumeCmd.PersistentFlags().StringArrayVar(&tagsInput, "tag", nil, "tags to identify resources to resume. \"key=value\"")
	ResumeCmd.PersistentFlags().BoolVar(&autoApprove, "auto-approve", false, "skip interactive approval")
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

		// get services
		services, err := lib.GetAWSServices(tags)
		if err != nil {
			fmt.Printf("Error getting services: %s\n", err)
			os.Exit(1)
		}

		// list services to resume
		fmt.Printf("Services to resume: \n\n")
		for _, service := range services {
			fmt.Println(service.GetARN())
		}

		// interactive confirmation
		fmt.Println()
		if !autoApprove {
			if _, err = confirmationPrompt.Run(); err != nil {
				fmt.Printf("No changes applied. %v\n", err)
				os.Exit(0)
			}
		}

		// resume services
		for _, service := range services {
			err := lib.Resume(service)
			if err != nil {
				fmt.Printf("Error resuming service: %s\n", service.GetARN())
				os.Exit(1)
			} else {
				fmt.Printf("Resuming service: %s\n", service.GetARN())
			}
		}
	},
}
