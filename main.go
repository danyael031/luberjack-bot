package main

import (
	"fmt"
	"image"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/micmonay/keybd_event"
)

func main() {
	//n := screenshot.NumActiveDisplays()
	// Left: 290x 350y
	// Right: 395x 350y
	var rightFlag bool = true
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	_ = rightFlag
	_ = kb
	// For linux, it is very important to wait 2 seconds
	fmt.Println("---------------Starting---------")
	time.Sleep(2 * time.Second)

	fmt.Println("---------------GO!!!---------")

	kb.SetKeys(keybd_event.VK_LEFT)
	kb.HasSuper(true)

	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	kb.Clear()

	// Game Bucle
	for i := 0; i <= 1000; i++ {
		time.Sleep(60 * time.Millisecond)

		if i >= 2 && i%2 == 0 {
			rightFlag = !verifyRisk()
		}

		if rightFlag {
			kb.SetKeys(keybd_event.VK_RIGHT)
		} else {
			kb.SetKeys(keybd_event.VK_LEFT)
		}

		err = kb.Launching()
		if err != nil {
			panic(err)
		}

		kb.Clear()

	}

}

func verifyRisk() bool {
	pixelLeft, err := screenshot.Capture(290, 345, 1, 1)
	if err != nil {
		panic(err)
	}

	pixelRigth, err := screenshot.Capture(395, 345, 1, 1)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5000; i++ {
		if (verifyRiskyPixel(pixelLeft)) || (verifyRiskyPixel(pixelRigth)) {
			break
		}
		//time.Sleep(8 * time.Millisecond)
		pixelLeft, err = screenshot.Capture(290, 345, 1, 1)
		if err != nil {
			panic(err)
		}

		pixelRigth, err = screenshot.Capture(395, 345, 1, 1)
		if err != nil {
			panic(err)
		}
	}

	if verifyRiskyPixel(pixelRigth) {
		fmt.Println("-----caution at right!!")
	}
	if verifyRiskyPixel(pixelLeft) {
		fmt.Println("-----caution at Left!!")
	}

	return verifyRiskyPixel(pixelRigth)

}

func verifyRiskyPixel(pixel *image.RGBA) bool {
	if int(pixel.Pix[0]) == 161 &&
		int(pixel.Pix[1]) == 116 &&
		int(pixel.Pix[2]) == 56 {
		return true
	}
	return false
}
