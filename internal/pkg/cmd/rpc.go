package cmd

import (
	"github.com/spf13/cobra"
)

// rpcCmd represents the rpc command
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "rpc接口服务",
	Run: func(cmd *cobra.Command, args []string) {
		if rpcCmdRun == nil {
			panic("cmd run not register")
		}
		rpcCmdRun(runArgs)
	},
}

func initRpcCmd() {
	rootCmd.AddCommand(rpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
