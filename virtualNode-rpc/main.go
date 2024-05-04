package main

import (
	"os"
	"virtualNode_rpc/server"
)

func main() {
	go server.InstanceInfoInform()
	code := server.Run(server.Cmd)
	os.Exit(code)
}
