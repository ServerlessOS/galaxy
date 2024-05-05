package main

import (
	"os"
	"scheduler_rpc/server"
)

func main() {
	go server.Schedule()
	code := server.Run(server.Cmd)
	os.Exit(code)
}
