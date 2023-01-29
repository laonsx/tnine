package cmd

import (
	"github.com/spf13/cobra"
)

// scriptCmd represents the script command
var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "脚本服务",
}

func initScriptCmd() {
	rootCmd.AddCommand(scriptCmd)
}
