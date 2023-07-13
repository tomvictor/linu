package main

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func createRefreshMenuItem(refreshAction chan bool) cocoa.NSMenuItem {
	refreshMenu := cocoa.NSMenuItem_New()
	refreshMenu.SetTitle("Refresh")
	refreshMenu.SetAction(objc.Sel("refreshAction:"))
	cocoa.DefaultDelegateClass.AddMethod("refreshAction:", func(_ objc.Object) {
		refreshAction <- true
	})
	return refreshMenu
}

func quitMenu() cocoa.NSMenuItem {
	itemQuit := cocoa.NSMenuItem_New()
	itemQuit.SetTitle("Quit")
	itemQuit.SetAction(objc.Sel("terminate:"))
	return itemQuit
}
