package main

import (
	"fmt"
	"os/exec"
)

const BaudRate115200 = "115200"
const BaudRate9600 = "9600"

var baudRate string = BaudRate115200

func openPort(port string) {
	fmt.Println("Opening port: ", port)
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf(`tell app "Terminal" to do script "minicom -b %s -D %s"`, baudRate, port))
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to run screen")
	}
}
