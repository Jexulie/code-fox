package commands

import (
	"code-fox/snippet"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
	"strconv"
)

/**
Executes snippets - cmd "execute"
*/

var executeCmd = &cobra.Command{
	Use:   "execute [id]",
	Short: "Execute a snippet by id",
	Long:  `Execute a snippet by id.`,
	Args:  cobra.MinimumNArgs(1),
	Aliases: []string{
		"e",
	},
	Example: "codefox execute 23",
	Run: func(cmd *cobra.Command, args []string) {
		idStrValue := args[0]
		idValue, err := strconv.ParseInt(idStrValue, 10, 64)
		if err != nil {
			fmt.Println("invalid id")
		}

		snip, err := snippet.GetById(idValue)
		if err != nil {
			fmt.Println("snippet not found")
			return
		}

		err = executeCommand(snip.Code)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func executeCommand(code string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", code)
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("sh", "-c", code)
	} else {
		cmd = exec.Command("sh", "-c", code)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
