package main

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func createSeparator() cocoa.NSMenuItem {
	return cocoa.NSMenuItem_Separator()
}

func createRefreshMenuItem(refreshAction chan bool) cocoa.NSMenuItem {
	refreshMenu := cocoa.NSMenuItem_New()
	refreshMenu.SetTitle("Refresh")
	refreshMenu.SetAction(objc.Sel("refreshAction:"))
	cocoa.DefaultDelegateClass.AddMethod("refreshAction:", func(_ objc.Object) {
		refreshAction <- true
	})
	return refreshMenu
}

func createQuitMenu() cocoa.NSMenuItem {
	quitMenu := cocoa.NSMenuItem_New()
	quitMenu.SetTitle("Quit")
	quitMenu.SetAction(objc.Sel("terminate:"))
	return quitMenu
}

func createSettingsMenu(baudRateSig chan string) cocoa.NSMenuItem {
	// Create a submenu
	baudRateMenu := cocoa.NSMenu_New()

	baudRateMenuItem := cocoa.NSMenuItem_New()
	baudRateMenuItem.SetTitle("Baud Rate")
	baudRateMenuItem.SetSubmenu(baudRateMenu)
	//mainMenu.AddItem(baudRateMenuItem)

	// Create child menu items for the submenu
	baudRate115200 := cocoa.NSMenuItem_New()
	baudRate115200.SetTitle(formatBaudRateTitle(BaudRate115200))
	baudRate115200.SetAction(objc.Sel("BaudRate115200:"))
	cocoa.DefaultDelegateClass.AddMethod("BaudRate115200:", func(_ objc.Object) {
		baudRateSig <- "115200"
	})
	//baudRate115200.SetAction(cocoa.GetSelector("childItem1Clicked:"))
	baudRate115200.SetTarget(cocoa.DefaultDelegate)
	baudRateMenu.AddItem(baudRate115200)

	baudRate9600 := cocoa.NSMenuItem_New()
	baudRate9600.SetTitle(formatBaudRateTitle(BaudRate9600))
	baudRate9600.SetAction(objc.Sel("baudRate9600:"))
	cocoa.DefaultDelegateClass.AddMethod("baudRate9600:", func(_ objc.Object) {
		baudRateSig <- "9600"
	})
	//baudRate9600.SetAction(cocoa.GetSelector("childItem2Clicked:"))
	baudRate9600.SetTarget(cocoa.DefaultDelegate)
	baudRateMenu.AddItem(baudRate9600)

	return baudRateMenuItem
}

func formatBaudRateTitle(title string) string {
	if baudRate == title {
		return fmt.Sprintf("☑ %s", title)
	}
	return fmt.Sprintf("☐ %s", title)
}
