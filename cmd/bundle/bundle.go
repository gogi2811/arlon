package bundle

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "bundle",
		Short:             "Manage configuration bundles",
		Long:              "Manage configuration bundles",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Usage()
		},
	}
	command.AddCommand(listBundlesCommand())
	command.AddCommand(dumpBundleCommand())
	command.AddCommand(createBundleCommand())
	command.AddCommand(deleteBundleCommand())
	command.AddCommand(updateBundleCommand())
	return command
}
