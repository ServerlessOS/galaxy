package main

import (
	"os"
	"scheduler_rpc/server"
)

func main() {
	code := server.Run(server.Cmd)
	os.Exit(code)
}
