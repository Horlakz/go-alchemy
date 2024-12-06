package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/struckchure/go-alchemy/components"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new component",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			categoryId  string
			componentId string
		)

		if len(args) > 0 {
			id := args[0]
			idSplit := strings.SplitN(id, ".", 2)
			categoryId = strings.SplitN(id, ".", 2)[0]

			if len(idSplit) > 1 {
				componentId = lo.Must(lo.Last(strings.SplitN(id, ".", 2)))
			}
		}

		if categoryId == "" {
			err := survey.AskOne(&survey.Select{
				Message: "Component Category",
				Options: components.ComponentCategoryOptions,
			}, &categoryId)
			if err != nil {
				color.Red("%s", err)
				return
			}
		}

		if componentId == "" {
			err := survey.AskOne(&survey.Select{
				Message: "Select Components",
				Options: components.ComponentMapping[categoryId],
			}, &componentId)
			if err != nil {
				color.Red("%s", err)
				return
			}
		}

		err := components.NewConfigService().Add(fmt.Sprintf("%s.%s", categoryId, componentId))
		if err != nil {
			color.Red("%s", err)
			return
		}
	},
}
