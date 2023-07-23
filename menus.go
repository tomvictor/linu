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
	quitMenu := cocoa.NSMenuItem_New()
	quitMenu.SetTitle("Quit")
	quitMenu.SetAction(objc.Sel("terminate:"))
	return quitMenu
}
