package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	println("Aqui")

	bitmap := robotgo.CaptureScreen(10, 20, 1024, 768)
	defer robotgo.FreeBitmap(bitmap)

	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap----- ", fx, fy)

	robotgo.SaveBitmap(bitmap, "testPrint.png")

}
