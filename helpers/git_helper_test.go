package helpers

import (
	"os/exec"
	"testing"
)

func TestCheckGitRepoType(t *testing.T) {
	// setup
	var cmd *exec.Cmd

	cmd = exec.Command("mkdir", "-p", "tmp/git-commit0")
	cmd.Run()
	cmd = exec.Command("mkdir", "-p", "tmp/git-commit2")
	cmd.Run()
	cmd = exec.Command("mkdir", "-p", "tmp/nongit")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit0", "init")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit2", "init")
	cmd.Run()
	cmd = exec.Command("touch", "tmp/git-commit2/test")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit2", "add", "test1")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit2", "commit", "-m", "test1")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit2", "add", "test2")
	cmd.Run()
	cmd = exec.Command("git", "-C", "tmp/git-commit2", "commit", "-m", "test2")
	cmd.Run()

	println("============")
	println(CheckGitRepoType("tmp/git-commit0"))
	println(CheckGitRepoType("tmp/git-commit2"))
	println(CheckGitRepoType("tmp/nongit"))
}
