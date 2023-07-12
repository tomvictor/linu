package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
	"runtime"
	"time"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {

	runtime.LockOSThread()

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		reloadClicked := make(chan bool)

		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		obj.Button().SetTitle("▶️ Ready")

		menu := cocoa.NSMenu_New()

		itemReload := cocoa.NSMenuItem_New()
		itemReload.SetTitle("Refresh")
		itemReload.SetAction(objc.Sel("reloadClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("reloadClicked:", func(_ objc.Object) {
			reloadClicked <- true
		})

		go func() {
			for {
				select {
				case <-time.After(60 * time.Second):
					fmt.Println("timer")
				case <-reloadClicked:
					fmt.Println("reloading")
					menu.RemoveAllItems()
					menu.AddItem(itemReload)
					ports, err := serial.GetPortsList()
					if err != nil {
						log.Fatal(err)
					}
					if len(ports) == 0 {
						log.Fatal("No serial ports found!")
					}
					for _, port := range ports {
						fmt.Printf("Found port: %v\n", port)
						portItem := cocoa.NSMenuItem_New()
						portItem.SetTitle(port)
						portItem.SetAction(objc.Sel(port))
						menu.AddItem(portItem)
					}

					itemQuit := cocoa.NSMenuItem_New()
					itemQuit.SetTitle("Quit")
					itemQuit.SetAction(objc.Sel("terminate:"))

					menu.AddItem(itemQuit)
					obj.SetMenu(menu)
				}
				labels := map[int]string{
					0: "🚀#%01d",
				}

				ports, _ := serial.GetPortsList()
				fmt.Println(len(ports))
				// updates to the ui should happen on the main thread to avoid strange bugs
				core.Dispatch(func() {
					obj.Button().SetTitle(fmt.Sprintf(labels[0], len(ports)))
				})
			}
		}()

		reloadClicked <- true
	})

	app.Run()
}
