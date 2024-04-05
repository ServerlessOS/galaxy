package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
)

var (
	command = *pflag.StringP("command", "c", "", "command for connect")
)

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine) // 绑定标志集
	url := "http://0.0.0.0:9999/ExecuteCommand"
	payload := []byte(viper.GetString("command"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
