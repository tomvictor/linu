package main

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type AppCtx struct {
	obj              cocoa.NSStatusItem
	refreshActionSig chan bool
	portActionSig    chan string
	mainMenu         cocoa.NSMenu
	refreshMenu      cocoa.NSMenuItem
}

func (app *AppCtx) initialize() {
	app.obj.Retain()
	app.obj.Button().SetTitle("🚀")
	app.refreshMenu = createRefreshMenuItem(app.refreshActionSig)
}

func (app *AppCtx) updateAppTitle(ports []string) {
	labels := map[int]string{
		0: "🚀#%01d",
	}
	// updates to the ui should happen on the main thread to avoid strange bugs
	core.Dispatch(func() {
		app.obj.Button().SetTitle(fmt.Sprintf(labels[0], len(ports)))
	})
}

func (app *AppCtx) reloadMainMenu(ports []string) {
	fmt.Println("reloading")
	app.mainMenu.RemoveAllItems()
	app.mainMenu.AddItem(app.refreshMenu)
	separatorItem := cocoa.NSMenuItem_Separator()
	app.mainMenu.AddItem(separatorItem)
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
		portItem := cocoa.NSMenuItem_New()
		portItem.SetTitle(port)
		portItem.SetAction(objc.Sel(fmt.Sprintf("%s:", port)))
		cocoa.DefaultDelegateClass.AddMethod(fmt.Sprintf("%s:", port), func(_ objc.Object) {
			app.portActionSig <- port
		})
		app.mainMenu.AddItem(portItem)
	}

	separatorItem2 := cocoa.NSMenuItem_Separator()
	app.mainMenu.AddItem(separatorItem2)

	app.mainMenu.AddItem(quitMenu())
	app.obj.SetMenu(app.mainMenu)
}
