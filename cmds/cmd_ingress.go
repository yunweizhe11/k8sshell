package cmds

import "github.com/spf13/cobra"

// Ingress 命令
var ingressRootCmd = cobra.Command{
	Use:     "ingress",
	Aliases: []string{"ing"},
	Short:   "ingress is used to manage kubernetes Ingresses",
}

// Ingress Create 命令
var ingressCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new ingress",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Ingress Update 命令
var ingressUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a ingress",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Ingress Get 命令
var ingressGetCmd = cobra.Command{
	Use:   "get",
	Short: "get ingress or ingress list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Ingress Delete 命令
var ingressDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a ingress",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
