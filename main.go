package main

import (
	"github.com/progrium/macdriver/cocoa"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	cocoa.TerminateAfterWindowsClose = false
	app := createCocoaApplication()
	app.Run()
}
