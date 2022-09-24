package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Up api",
	Long:  `Up backend api `,
	Run:   apiCobraCommand,
}

func apiCobraCommand(cmd *cobra.Command, args []string) {
	exec.Command("pwd")
	teste := exec.Command("pwd")
	if err := teste.Run(); err != nil {
		panic(err)
	}

}
func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
