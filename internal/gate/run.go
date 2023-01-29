package gate

import (
	"log"

	"github.com/laonsx/tnine/internal/pkg/cmd"
)

func init() {
	cmd.RegisterApi(runApi)
	cmd.RegisterRpc(runRpc)
}

func runApi(args cmd.RunArgs) {
	log.Println("gate api run")
}

func runRpc(args cmd.RunArgs) {
	log.Println("gate rpc run")
}
