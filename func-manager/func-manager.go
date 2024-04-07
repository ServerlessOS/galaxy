package main

import (
	"func-manager/cmd"
	"os"
)

func main() {
	code := cmd.Run(cmd.Cmd)
	os.Exit(code)
}
