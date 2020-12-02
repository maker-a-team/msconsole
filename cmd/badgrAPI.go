package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tempor1s/msconsole/modules"
)

func init() {
	rootCmd.AddCommand(badgrCommand)
}

var badgrCommand = &cobra.Command{
	Use:   "badgr [badgr password]",
	Short: "Get Access Token for Badgr.",
	Long:  "This command will allow the user to get an access token for Badgr.",
	Run: func(cmd *cobra.Command, args []string) {
		modules.BadgrModule(cmd, args)
	},
}
