package main

import (
	"fmt"

	"github.com/rhomari/systrayicon"
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
