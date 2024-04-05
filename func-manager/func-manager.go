package main

import (
	"func-manager/cmd"
	"os"
)

func main() {
	command := cmd.NewGatewayCommand()
	code := cmd.Run(command)
	os.Exit(code)
}
