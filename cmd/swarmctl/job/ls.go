package job

import "github.com/spf13/cobra"

var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List jobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)