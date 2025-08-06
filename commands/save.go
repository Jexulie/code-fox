package commands

import (
	"code-fox/file"
	"code-fox/snippet"
	"code-fox/tag"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

/*
*
Saves snippets - cmd "save"
Flags
--lang=go
--code="console.log"
--tag=network,utility
--description="Useful for something"
--file=retry.go
*/
var languageValue string
var tagValue string
var descriptionValue string
var filePathValue string
var codeValue string

var saveCmd = &cobra.Command{
	Use:   "save [title]",
	Short: "Create a code snippet",
	Long:  `save a code snippet or a file relative to current directory.`,
	Args:  cobra.ExactArgs(1),
	Aliases: []string{
		"s",
	},
	Example: "codefox save \"Curl get my-site\" -l=bash -t=curl,command -c=\"curl -L my-site.com\" -d=\"Curl get command to my-site\"",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// validate either code or file provided
		if codeValue == "" && filePathValue == "" {
			return fmt.Errorf("either --file (-f) or --code (-c) must be provided\n")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		titleValue := args[0]

		// if it's a file get contents
		if filePathValue != "" {
			getCodeFromFile()
		}

		s := snippet.NewSnippet(
			titleValue,
			codeValue,
			languageValue,
			descriptionValue,
		)

		err := s.Create()
		if err != nil {
			fmt.Println(err)
			return
		}

		if tagValue != "" {
			tagIds, err := getOrSetTags()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = saveSnippetTagRelation(s.Id, tagIds)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("Snippet saved successfully")
	},
}

func addSaveCmdFlags() {
	saveCmd.Flags().StringVarP(&languageValue, "lang", "l", "", "The language of the snippet")
	saveCmd.Flags().StringVarP(&tagValue, "tag", "t", "", "The tag of the snippet")
	saveCmd.Flags().StringVarP(&descriptionValue, "desc", "d", "", "The description of the snippet")
	saveCmd.Flags().StringVarP(&filePathValue, "file", "f", "", "The file name of the snippet with it's extension")
	saveCmd.Flags().StringVarP(&codeValue, "code", "c", "", "The code of the snippet")
	err := saveCmd.MarkFlagRequired("lang")

	if err != nil {
		log.Fatal(err)
	}

}

func getCodeFromFile() {
	contents, err := file.GetFileContents(filePathValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	codeValue = string(contents)
}

func getOrSetTags() ([]int64, error) {
	var tagIds []int64

	splitTags := strings.Split(tagValue, ",")

	for _, i := range splitTags {
		var tagId int64

		foundTag, err := tag.GetTagByName(i)
		if err != nil {
			t := tag.NewTag(i)
			err = t.Create()
			if err != nil {
				return tagIds, err
			}
			tagId = t.Id
		} else {
			tagId = foundTag.Id
		}

		tagIds = append(tagIds, tagId)
	}

	return tagIds, nil
}

func saveSnippetTagRelation(snippetId int64, tagIds []int64) error {
	for _, tagId := range tagIds {
		err := tag.AddSnippetTag(snippetId, tagId)
		if err != nil {
			return err
		}
	}

	return nil
}
