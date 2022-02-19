package screens

import (
	"github.com/progrium/macdriver/cocoa"
)

func Active() string {
	return cocoa.NSScreen_Main().LocalizedName().String()
}

func IsActive(desired string) bool {
	if desired == cocoa.NSScreen_Main().LocalizedName().String() {
		return true
	}
	return false
}

func List() []string {
	screens := cocoa.NSScreen_Screens()
	var screenNames []string
	for _, screen := range screens {
		screenNames = append(screenNames, screen.LocalizedName().String())
	}
	return screenNames
}