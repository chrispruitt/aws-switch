package cmd

import (
	"fmt"
	"os"

	"github.com/chrispruitt/aws-switch/lib"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	tagsInput          []string
	autoApprove        bool
	confirmationPrompt = promptui.Prompt{
		Label:     "Would you like to apply these changes",
		IsConfirm: true,
	}
)

func init() {
	HaltCmd.PersistentFlags().StringArrayVar(&tagsInput, "tag", nil, "tags to identify resources to halt - \"key=value\"")
	HaltCmd.PersistentFlags().BoolVar(&autoApprove, "auto-approve", false, "skip interactive approval")
}

var HaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Halt an aws service",
	Run: func(cmd *cobra.Command, args []string) {

		if len(tagsInput) == 0 {
			fmt.Println("--tags required")
			os.Exit(1)
		}

		// validate tags
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

		// list services to halt
		fmt.Printf("Services to halt: \n\n")
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
		// halt services
		for _, service := range services {
			err := lib.Halt(service)
			if err != nil {
				fmt.Printf("Error halting service: %s\n", service.GetARN())
				os.Exit(1)
			} else {
				fmt.Printf("Halting service: %s\n", service.GetARN())
			}
		}
	},
}
