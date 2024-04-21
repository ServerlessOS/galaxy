package main

import (
	"dispatcher_rpc/server"
	"os"
)

func main() {
	code := server.Run(server.Cmd)
	os.Exit(code)
}
