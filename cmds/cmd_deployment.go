package cmds

import "github.com/spf13/cobra"

// Deployment 命令
var deploymentRootCmd = cobra.Command{
	Use:     "deployment",
	Aliases: []string{"deploy"},
	Short:   "deployment is used to manage kubernetes Deployments",
}

// Deployment Create 命令
var deploymentCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new deployment",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Deployment Update 命令
var deploymentUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a deployment",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Deployment Get 命令
var deploymentGetCmd = cobra.Command{
	Use:   "get",
	Short: "get deployment or deployment list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Deployment Delete 命令
var deploymentDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a deployment",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
