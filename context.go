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
	baudRateSig      chan string
	mainMenu         cocoa.NSMenu
	refreshMenu      cocoa.NSMenuItem
	settingsMenu     cocoa.NSMenuItem
	quitMenu         cocoa.NSMenuItem
}

func (app *AppCtx) initialize() {
	app.obj.Retain()
	app.obj.Button().SetTitle("🚀")
	app.refreshMenu = createRefreshMenuItem(app.refreshActionSig)
	app.settingsMenu = createSettingsMenu(app.baudRateSig)
	app.quitMenu = createQuitMenu()
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
	app.mainMenu.AddItem(createSeparator())
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

	app.mainMenu.AddItem(createSeparator())
	app.mainMenu.AddItem(app.settingsMenu)
	updateSettingsMenu(app.settingsMenu, app.baudRateSig)
	app.mainMenu.AddItem(createSeparator())
	app.mainMenu.AddItem(app.quitMenu)
	app.obj.SetMenu(app.mainMenu)
}
