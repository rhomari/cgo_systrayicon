package main

import (
	"Systray-Icon/systrayicon/systrayicon"
	"fmt"
)

func main() {
	systrayicon.OnMouseClick = func(lparm int) {
		switch lparm {
		case systrayicon.WM_LBUTTONUP:
			fmt.Println("Left mouse button click")
			break
		case systrayicon.WM_RBUTTONUP:
			fmt.Println("Right mouse button click")
			break

		}
	}
	systrayicon.CreateSystrayIcon()

}
