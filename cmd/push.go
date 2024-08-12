package cmd

import (
	"fmt"
	"php-git-tool/internal/git"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote refs along with associated objects",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := git.Push()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(output)
	},
}
