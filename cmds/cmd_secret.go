package cmds

import (
	"github.com/spf13/cobra"
)

// secret 命令
var secretRootCmd = cobra.Command{
	Use:   "secret",
	Short: "secret is used to manage kubernetes secrets",
}

// secret Create 命令
var secretCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new secret",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// secret Update 命令
var secretUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a secret",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// secret Get 命令
var secretGetCmd = cobra.Command{
	Use:   "get",
	Short: "get secret or secret list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// secret Delete 命令
var secretDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a secret",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
