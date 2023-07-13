package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
	"os/exec"
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
		portClicked := make(chan string)

		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		obj.Button().SetTitle("🚀")

		menu := cocoa.NSMenu_New()

		itemReload := cocoa.NSMenuItem_New()
		itemReload.SetTitle("Refresh")
		itemReload.SetAction(objc.Sel("reloadClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("reloadClicked:", func(_ objc.Object) {
			reloadClicked <- true
		})

		sub := cocoa.NSMenu_New()
		sub.SetTitle("subR")

		sub1 := cocoa.NSMenuItem_New()
		sub1.SetTitle("sub1")

		sub2 := cocoa.NSMenuItem_New()
		sub2.SetTitle("sub1")

		sub.AddItem(sub1)
		sub.AddItem(sub2)

		itemReload.SetSubmenu(sub)

		go func() {
			for {
				select {
				case <-time.After(60 * time.Second):
					fmt.Println("timer")
				case port := <-portClicked:
					fmt.Println(port, " clicked")
					cmd := exec.Command("osascript", "-e", fmt.Sprintf(`tell app "Terminal" to do script "screen %s 115200"`, port))
					// Run the command
					err := cmd.Run()
					if err != nil {
						fmt.Println("failed to run cmd")
					}
				case <-reloadClicked:
					fmt.Println("reloading")
					menu.RemoveAllItems()
					menu.AddItem(itemReload)
					separatorItem := cocoa.NSMenuItem_Separator()
					menu.AddItem(separatorItem)
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
						portItem.SetAction(objc.Sel(fmt.Sprintf("%s:", port)))
						cocoa.DefaultDelegateClass.AddMethod(fmt.Sprintf("%s:", port), func(_ objc.Object) {
							portClicked <- port
						})
						menu.AddItem(portItem)
					}

					itemQuit := cocoa.NSMenuItem_New()
					itemQuit.SetTitle("Quit")
					itemQuit.SetAction(objc.Sel("terminate:"))

					separatorItem2 := cocoa.NSMenuItem_Separator()
					menu.AddItem(separatorItem2)

					menu.AddItem(itemQuit)
					obj.SetMenu(menu)
				}
				labels := map[int]string{
					0: "🚀#%01d",
				}

				ports, _ := serial.GetPortsList()
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
