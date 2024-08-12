package cmd

import (
	"fmt"
	"php-git-tool/internal/git"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func formatStatusOutput(output string) string {
	lines := strings.Split(output, "\n")
	var formattedOutput strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "On branch") {
			formattedOutput.WriteString(fmt.Sprintf("%s %s\n", color.YellowString("Branch:"), color.CyanString(line[10:])))
		} else if strings.HasPrefix(line, "Changes to be committed:") {
			formattedOutput.WriteString(color.GreenString("\nChanges to be committed:\n"))
		} else if strings.HasPrefix(line, "Changes not staged for commit:") {
			formattedOutput.WriteString(color.RedString("\nChanges not staged for commit:\n"))
		} else if strings.HasPrefix(line, "Untracked files:") {
			formattedOutput.WriteString(color.BlueString("\nUntracked files:\n"))
		} else if strings.HasPrefix(line, "modified:") || strings.HasPrefix(line, "new file:") || strings.HasPrefix(line, "deleted:") {
			formattedOutput.WriteString(fmt.Sprintf("  %s\n", line))
		} else {
			formattedOutput.WriteString(fmt.Sprintf("%s\n", line))
		}
	}

	return formattedOutput.String()
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := git.Status()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(output)
	},
}
