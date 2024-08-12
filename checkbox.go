// package main

// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"strings"

// 	"github.com/AlecAivazis/survey/v2"
// )

// func main() {
// 	// Lấy danh sách các branch local
// 	branches, err := getLocalBranches()
// 	if err != nil {
// 		log.Fatalf("Failed to get local branches: %v", err)
// 	}

// 	// Tạo câu hỏi checkbox với các branch
// 	var selectedBranches []string
// 	err = survey.AskOne(&survey.MultiSelect{
// 		Message: "Chọn các branch:",
// 		Options: branches,
// 	}, &selectedBranches)
// 	if err != nil {
// 		log.Fatalf("Failed to get user input: %v", err)
// 	}

// 	// Hiển thị các branch đã chọn
// 	fmt.Println("Các branch đã chọn:")
// 	for _, branch := range selectedBranches {
// 		fmt.Printf("- %s\n", branch)
// 	}
// }

// // Lấy danh sách các branch local
// func getLocalBranches() ([]string, error) {
// 	cmd := exec.Command("git", "branch")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Chuyển đổi output thành danh sách các branch
// 	lines := strings.Split(string(output), "\n")
// 	var branches []string
// 	for _, line := range lines {
// 		branch := strings.TrimSpace(line)
// 		if branch != "" {
// 			// Xóa dấu * nếu có
// 			branch = strings.TrimPrefix(branch, "* ")
// 			branches = append(branches, branch)
// 		}
// 	}

// 	return branches, nil
// }
