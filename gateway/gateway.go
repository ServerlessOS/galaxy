package main

import (
	"gateway/cmd"
	"os"
)

func main() {
	code := cmd.Run(cmd.Cmd)
	os.Exit(code)
}
