package main

import (
	"fmt"
	"os/exec"
)

var BaudRate string = "115200"

func openPort(port string) {
	fmt.Println("Opening port: ", port)
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf(`tell app "Terminal" to do script "minicom -b %s -D %s"`, BaudRate, port))
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to run screen")
	}
}
