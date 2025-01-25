package gen_config

import "github.com/spf13/cobra"

func NewSetcloneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "bulk-clone",
		Short:        "bulk-clone subcommand to manage GitHub org repositories",
		SilenceUsage: true,
	}

	cmd.AddCommand(newSetcloneGiteaCmd())
	cmd.AddCommand(newSetcloneGithubCmd())
	cmd.AddCommand(newSetcloneGitlabCmd())

	return cmd
}
