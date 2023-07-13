package main

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

func applicationCallback(n objc.Object) {

	refreshAction := make(chan bool)
	portAction := make(chan string)

	obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
	obj.Retain()
	obj.Button().SetTitle("🚀")

	go applicationControlPanel(portAction, refreshAction, obj)

	refreshAction <- true
}

func createCocoaApplication() cocoa.NSApplication {
	app := cocoa.NSApp_WithDidLaunch(applicationCallback)
	return app
}
