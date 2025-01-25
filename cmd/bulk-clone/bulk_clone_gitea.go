package bulk_clone

import (
	"fmt"
	"github.com/spf13/cobra"
)

type bulkCloneGiteaOptions struct {
	targetPath string
	orgName    string
}

func defaultBulkCloneGiteaOptions() *bulkCloneGiteaOptions {
	return &bulkCloneGiteaOptions{}
}

func newBulkCloneGiteaCmd() *cobra.Command {
	o := defaultBulkCloneGiteaOptions()

	cmd := &cobra.Command{
		Use:   "gitea",
		Short: "Clone repositories from a Gitea organization",
		Args:  cobra.NoArgs,
		RunE:  o.run,
	}

	cmd.Flags().StringVarP(&o.targetPath, "targetPath", "t", o.targetPath, "targetPath")
	cmd.Flags().StringVarP(&o.orgName, "orgName", "o", o.orgName, "orgName")

	return cmd
}

func (o *bulkCloneGiteaOptions) run(_ *cobra.Command, args []string) error {
	if o.targetPath == "" || o.orgName == "" {
		return fmt.Errorf("both targetPath and orgName must be specified")
	}

	//err := gitea_org.RefreshAll(o.targetPath, o.orgName)
	//if err != nil {
	//	return err
	//}

	return nil
}
