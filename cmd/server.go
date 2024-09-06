package cmd

import (
	v1 "github.com/nilssonr/agentside/api/v1"
	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/interaction"
	"github.com/nilssonr/agentside/logging"
	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/repository/postgres"
	"github.com/nilssonr/agentside/skill"
	"github.com/nilssonr/agentside/tenant"
	"github.com/nilssonr/agentside/user"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			logging.Module,
			postgres.Module,
			customer.Module,
			interaction.Module,
			queue.Module,
			skill.Module,
			tenant.Module,
			user.Module,
			v1.Module,
		)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCommand)
}
