package cmd

import (
	"fmt"
	"php-git-tool/internal/git"

	"github.com/spf13/cobra"
)

var commitMsg string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := git.Commit(commitMsg)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(output)
	},
}

func init() {
	commitCmd.Flags().StringVarP(&commitMsg, "message", "m", "", "Commit message")
	_ = commitCmd.MarkFlagRequired("message")
}
