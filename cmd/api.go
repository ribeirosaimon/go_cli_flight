package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Up api",
	Long:  `Up backend api `,
	Run:   apiCobraCommand,
}

//https://github.com/docker/compose/blob/9d73cc88ccfafbd3cb1bb83287f9e9e2d686d393/cmd/compose/up.go#L111

func apiCobraCommand(cobraCommand *cobra.Command, args []string) {

	cmd, err := exec.Command("/bin/bash", "docker-compose", "up").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)

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
