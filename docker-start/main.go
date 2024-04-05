package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

var (
	port = *pflag.IntP("port", "p", 9999, "port for connect,such as 9999")
)

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine) // 绑定标志集
	http.HandleFunc("/ExecuteCommand", ExecuteCommand)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
func ExecuteCommand(w http.ResponseWriter, req *http.Request) {
	reader := bufio.NewReader(req.Body)
	command, _ := reader.ReadString('\n')
	cmd := exec.Command("docker", command)
	fmt.Println("Execute command:", "docker", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("command out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	w.Write(out)
	return
}
