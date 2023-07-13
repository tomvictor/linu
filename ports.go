package main

import (
	"fmt"
	"os/exec"
)

func openPort(port string) {
	fmt.Println("Opening port: ", port)
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf(`tell app "Terminal" to do script "screen %s 115200"`, port))
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to run screen")
	}
}
