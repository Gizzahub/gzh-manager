package bulk_clone

import (
	"fmt"
	gitlabpkg "github.com/gizzahub/gzh-manager-go/pkg/gitlab"
	"github.com/spf13/cobra"
)

type bulkCloneGitlabOptions struct {
	targetPath  string
	groupName   string
	recursively bool
}

func defaultBulkCloneGitlabOptions() *bulkCloneGitlabOptions {
	return &bulkCloneGitlabOptions{}
}

func newBulkCloneGitlabCmd() *cobra.Command {
	o := defaultBulkCloneGitlabOptions()

	cmd := &cobra.Command{
		Use:   "gitlab",
		Short: "Clone repositories from a GitLab group",
		Args:  cobra.NoArgs,
		RunE:  o.run,
	}

	cmd.Flags().StringVarP(&o.targetPath, "targetPath", "t", o.targetPath, "targetPath")
	cmd.Flags().StringVarP(&o.groupName, "groupName", "g", o.groupName, "groupName")
	cmd.Flags().BoolVarP(&o.recursively, "recursively", "r", o.recursively, "recursively")

	return cmd
}

func (o *bulkCloneGitlabOptions) run(_ *cobra.Command, args []string) error {
	if o.targetPath == "" || o.groupName == "" {
		return fmt.Errorf("both targetPath and groupName must be specified")
	}

	err := gitlabpkg.RefreshAll(o.targetPath, o.groupName)
	if err != nil {
		return err
	}

	return nil
}
