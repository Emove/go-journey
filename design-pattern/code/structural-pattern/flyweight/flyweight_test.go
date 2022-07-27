package flyweight

import "testing"

func TestImageViewer(t *testing.T) {
	viewer := NewImageViewer("image.png")
	viewer.Display()
}

func TestFlyweight(t *testing.T) {
	viewer := NewImageViewer("image.png")
	imageViewer := NewImageViewer("image.png")
	if viewer.ImageFlyweight != imageViewer.ImageFlyweight {
		t.Fail()
	}
}
