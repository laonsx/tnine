package cmd

import (
	"github.com/spf13/cobra"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务服务",
	Run: func(cmd *cobra.Command, args []string) {
		if cronCmdRun == nil {
			panic("cmd run not register")
		}
		cronCmdRun(runArgs)
	},
}

func initCronCmd() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
