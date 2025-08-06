package commands

import (
	"code-fox/helpers"
	"code-fox/snippet"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

/*
**
Lists snippets - cmd "list"
*/
const (
	idWidth          = 4
	titleWidth       = 25
	langWidth        = 10
	tagsWidth        = 15
	descriptionWidth = 15
	codeWidth        = 12
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Long:  "List all snippets",
	Run: func(cmd *cobra.Command, args []string) {

		snippets, err := snippet.GetSnippets(tagValue)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found")
			return
		}

		// Header
		fmt.Printf("%-*s %-*s %-*s %-*s %-*s %-*s\n",
			idWidth, "ID",
			titleWidth, "Title",
			langWidth, "Language",
			tagsWidth, "Tags",
			descriptionWidth, "Description",
			codeWidth, "Code")

		fmt.Println(strings.Repeat("-", idWidth+titleWidth+langWidth+tagsWidth+descriptionWidth+codeWidth+4))

		for _, snippet := range snippets {
			formattedCode := helpers.RemoveSpecialChars(helpers.LimitString(snippet.Code, codeWidth))
			fmt.Printf("%-*d %-*s %-*s %-*s %-*s %-*s\n",
				idWidth, snippet.Id,
				titleWidth, helpers.LimitString(snippet.Title, titleWidth),
				langWidth, helpers.LimitString(snippet.Language, langWidth),
				tagsWidth, helpers.LimitString(snippet.Tags, tagsWidth),
				descriptionWidth, helpers.LimitString(snippet.Description, descriptionWidth),
				codeWidth, formattedCode)
		}
	},
}

func addListCmdFlags() {
	listCmd.Flags().StringVarP(&tagValue, "tag", "t", "", "list by tag")
}
