package example

import (
	"log"

	"github.com/laonsx/tnine/internal/example/internal/service"
	"github.com/laonsx/tnine/internal/pkg/cmd"
)

func init() {
	cmd.RegisterApi(runApi)
	cmd.RegisterRpc(runRpc)
}

func runApi(args cmd.RunArgs) {

	log.Println("example api run")

	log.Println("log err from down stream")
	_ = service.LogErr(true)
	log.Println("log err from up stream")
	err := service.LogErr(false)
	if err != nil {
		log.Printf("%+v", err)
	}
}

func runRpc(args cmd.RunArgs) {
	log.Println("example rpc run")
}
