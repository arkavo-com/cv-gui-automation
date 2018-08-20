package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
)

func main() {
	const displayIndex = 0
	if screenshot.NumActiveDisplays() > 0 {
		window := gocv.NewWindow(fmt.Sprintf("Capture window %v", displayIndex))
		var mat gocv.Mat
		var err error
		for {
			mat, err = MatFromVideoCaptureDisplay(displayIndex)
			if err != nil {
				break
			}
			if mat.Empty() {
				continue
			}
			// TODO add TensorFlow Classifier
			robotgo.ScrollMouse(2, "up")
			window.IMShow(mat)
			if window.WaitKey(1) == 27 {
				break
			}
		}
		mat.Close()
		window.Close()
	}
}

func MatFromVideoCaptureDisplay(displayIndex int) (gocv.Mat, error) {
	var mat gocv.Mat
	var err error
	// TODO review robotgo screen capture
	capture, err := screenshot.CaptureDisplay(displayIndex)
	if err != nil {
		return mat, err
	}
	mat, err = gocv.ImageToMatRGBA(capture)
	return mat, err
}
