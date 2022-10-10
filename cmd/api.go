package cmd

import (
	"github.com/spf13/cobra"
	"github/ribeirosaimon/go_flight_cli/until"
	"os/exec"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Up api",
	Long:  `Up backend api `,
	Run:   apiCobraCommand,
}

func apiCobraCommand(cmdCobra *cobra.Command, args []string) {
	commands := exec.Command("docker-compose", "up", "-d")
	until.Execute(commands)
}
func init() {
	rootCmd.AddCommand(apiCmd)
}
