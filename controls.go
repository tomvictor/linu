package main

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"go.bug.st/serial"
	"log"
	"time"
)

func applicationControlPanel(portAction chan string, refreshAction chan bool, obj cocoa.NSStatusItem) {
	mainMenu := cocoa.NSMenu_New()

	refreshMenu := createRefreshMenuItem(refreshAction)

	for {
		ports, _ := serial.GetPortsList()
		if len(ports) == 0 {
			log.Fatal("No serial ports found!")
		}

		select {
		case <-time.After(60 * time.Second):
			fmt.Println("Timeout")
		case port := <-portAction:
			openPort(port)
		case <-refreshAction:
			reloadMainMenu(portAction, mainMenu, refreshMenu, ports, obj)
		}
		labels := map[int]string{
			0: "🚀#%01d",
		}

		// updates to the ui should happen on the main thread to avoid strange bugs
		core.Dispatch(func() {
			obj.Button().SetTitle(fmt.Sprintf(labels[0], len(ports)))
		})
	}
}

func reloadMainMenu(portAction chan string, mainMenu cocoa.NSMenu, refreshMenu cocoa.NSMenuItem, ports []string, obj cocoa.NSStatusItem) {
	fmt.Println("reloading")
	mainMenu.RemoveAllItems()
	mainMenu.AddItem(refreshMenu)
	separatorItem := cocoa.NSMenuItem_Separator()
	mainMenu.AddItem(separatorItem)
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
		portItem := cocoa.NSMenuItem_New()
		portItem.SetTitle(port)
		portItem.SetAction(objc.Sel(fmt.Sprintf("%s:", port)))
		cocoa.DefaultDelegateClass.AddMethod(fmt.Sprintf("%s:", port), func(_ objc.Object) {
			portAction <- port
		})
		mainMenu.AddItem(portItem)
	}

	separatorItem2 := cocoa.NSMenuItem_Separator()
	mainMenu.AddItem(separatorItem2)

	mainMenu.AddItem(quitMenu())
	obj.SetMenu(mainMenu)
}
