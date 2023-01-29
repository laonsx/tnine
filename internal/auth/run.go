package auth

import (
	"log"

	"github.com/laonsx/tnine/internal/pkg/cmd"
)

func init() {
	cmd.RegisterApi(runApi)
	cmd.RegisterRpc(runRpc)
}

func runApi(args cmd.RunArgs) {
	log.Println("auth api run")
}

func runRpc(args cmd.RunArgs) {
	log.Println("auth rpc run")
}
