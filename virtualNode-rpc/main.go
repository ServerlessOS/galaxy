package main

import (
	"os"
	"virtualNode_rpc/server"
)

func main() {
	code := server.Run(server.Cmd)
	os.Exit(code)
}
