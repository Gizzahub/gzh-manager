package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckGitRepoType(dir string) (string, error) {
	// Git 디렉토리인지 확인
	_, err := os.Stat(dir + "/.git")
	if os.IsNotExist(err) {
		return "none", nil
	}
	//cmd := exec.Command("git", "-C", dir, "rev-parse", "--is-inside-work-tree")
	//err := cmd.Run()
	//output, err := cmd.Output()
	//if err != nil {
	//	return "", fmt.Errorf("not a git repository: %v", err)
	//}

	// Git 커밋 수 확인
	cmd := exec.Command("git", "-C", dir, "rev-list", "--count", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		fmt.Errorf("failed to count commits: %v", err)
		panic(err)
	}

	commitCount := strings.TrimSpace(string(output))
	println(commitCount)

	// 0 커밋인 경우
	if commitCount == "0" {
		return "empty", nil
	}

	// 커밋이 있는 경우
	return "normal", nil
}
