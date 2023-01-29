package cmd

import (
	"github.com/spf13/cobra"
)

// scriptRpcCmd represents the scriptRpc command
var scriptRpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "脚本运行api接口",
	Run: func(cmd *cobra.Command, args []string) {
		if scriptRpcCmdRun == nil {
			panic("cmd run not register")
		}
		scriptRpcCmdRun(runArgs)
	},
}

func initScriptRpcCmd() {
	rootCmd.AddCommand(scriptRpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolsRpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolsRpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
