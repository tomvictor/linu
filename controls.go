package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
	"time"
)

func applicationControlPanel(appCtx AppCtx) {

	for {
		ports, _ := serial.GetPortsList()
		if len(ports) == 0 {
			log.Fatal("No serial ports found!")
		}

		select {
		case <-time.After(60 * time.Second):
			fmt.Println("Timeout")
		case port := <-appCtx.portActionSig:
			openPort(port)
		case <-appCtx.refreshActionSig:
			appCtx.reloadMainMenu(ports)
		case baudSel := <-appCtx.baudRateSig:
			baudRate = baudSel
			fmt.Println(baudSel)
			appCtx.reloadMainMenu(ports)
		}
		appCtx.updateAppTitle(ports)
	}
}
