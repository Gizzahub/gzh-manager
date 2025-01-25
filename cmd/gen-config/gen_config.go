package gen_config

import "github.com/spf13/cobra"

func NewGenConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gen-config",
		Short:        "gen-config subcommand to manage GitHub org repositories",
		SilenceUsage: true,
	}

	cmd.AddCommand(newGenConfigGithubCmd())

	return cmd
}
