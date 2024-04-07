package main

import (
	"func-manager/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	code := cmd.Run(cmd.Cmd)
	go cmd.GracefulExit()
	<-cmd.ExitCh
	log.Printf("Server exit")
	os.Exit(code)
}
