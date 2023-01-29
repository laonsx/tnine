package cmd

import (
	"github.com/spf13/cobra"
)

// scriptApiCmd represents the scriptApi command
var scriptApiCmd = &cobra.Command{
	Use:   "api",
	Short: "脚本运行api接口",
	Run: func(cmd *cobra.Command, args []string) {
		if scriptApiCmdRun == nil {
			panic("cmd run not register")
		}
		scriptApiCmdRun(runArgs)
	},
}

func initScriptApiCmd() {
	scriptCmd.AddCommand(scriptApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolsApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolsApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
