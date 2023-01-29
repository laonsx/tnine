package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "[service]",
	Short: "后端程序",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

type Run func(RunArgs)

type RunArgs struct {
}

var apiCmdRun, consumerCmdRun, cronCmdRun, rpcCmdRun, scriptApiCmdRun, scriptRpcCmdRun Run

var runArgs RunArgs

func RegisterApi(run Run) {
	initApiCmd()
	apiCmdRun = run
}

func RegisterConsumer(run Run) {
	initConsumerCmd()
	consumerCmdRun = run
}

func RegisterCron(run Run) {
	initCronCmd()
	cronCmdRun = run
}

func RegisterRpc(run Run) {
	initRpcCmd()
	rpcCmdRun = run
}

func RegisterScriptApi(run Run) {
	initScriptApiCmd()
	scriptApiCmdRun = run
}

func RegisterScriptRpc(run Run) {
	initScriptRpcCmd()
	scriptRpcCmdRun = run
}
