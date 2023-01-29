package cmd

import (
	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "异步消费者服务",
	Run: func(cmd *cobra.Command, args []string) {
		if consumerCmdRun == nil {
			panic("cmd run not register")
		}
		consumerCmdRun(runArgs)
	},
}

func initConsumerCmd() {
	rootCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
