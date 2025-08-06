package commands

import (
	"code-fox/snippet"
	"code-fox/tag"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

/**
Edits snippet - cmd "edit"
*/

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edits snippet by id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStrValue := args[0]
		idValue, err := strconv.ParseInt(idStrValue, 10, 64)
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		snip, err := snippet.GetById(idValue)
		if err != nil {
			fmt.Println("snippet not found")
			return
		}

		if codeValue == "" && filePathValue != "" {
			getCodeFromFile()
		}

		if languageValue != "" {
			snip.Language = languageValue
		}

		if tagValue != "" {
			err := tag.DeleteAllTagRelations(snip.Id)
			if err != nil {
				fmt.Println(err)
				return
			}

			tagIds, err := getOrSetTags()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = saveSnippetTagRelation(snip.Id, tagIds)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if descriptionValue != "" {
			snip.Description = descriptionValue
		}

		if codeValue != "" {
			snip.Code = codeValue
		}

		err = snip.Update()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("snippet updated")
	},
}

func addEditCmdFlags() {
	editCmd.Flags().StringVarP(&languageValue, "lang", "l", "", "The language of the snippet")
	editCmd.Flags().StringVarP(&tagValue, "tag", "t", "", "The tag of the snippet")
	editCmd.Flags().StringVarP(&descriptionValue, "desc", "d", "", "The description of the snippet")
	editCmd.Flags().StringVarP(&filePathValue, "file", "f", "", "The file name of the snippet with it's extension")
	editCmd.Flags().StringVarP(&codeValue, "code", "c", "", "The code of the snippet")
}
