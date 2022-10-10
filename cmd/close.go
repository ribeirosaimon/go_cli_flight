package cmd

import (
	"github.com/spf13/cobra"
	"github/ribeirosaimon/go_flight_cli/until"
	"os/exec"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Down api",
	Long:  `Down backend api `,
	Run:   closeCobraCommand,
}

func closeCobraCommand(cmd *cobra.Command, args []string) {
	commands := [2]*exec.Cmd{}
	commands[0] = exec.Command("docker-compose", "down")
	commands[1] = exec.Command("/bin/bash", "cmd/scripts/docker-prune.sh")

	for _, command := range commands {
		until.Execute(command)
	}
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
