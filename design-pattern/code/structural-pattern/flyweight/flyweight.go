package flyweight

import "fmt"

var imageFactory *ImageFlyweightFactory

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

type ImageFlyweight struct {
	data string
}

type ImageViewer struct {
	*ImageFlyweight
}

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (factory *ImageFlyweightFactory) Get(imageName string) *ImageFlyweight {
	image := factory.maps[imageName]
	if image == nil {
		image = NewImageFlyweight(imageName)
		factory.maps[imageName] = image
	}
	return image
}

func NewImageFlyweight(imageName string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", imageName)
	return &ImageFlyweight{
		data: data,
	}
}

func (image *ImageFlyweight) Data() string {
	return image.data
}

func NewImageViewer(imageName string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(imageName)
	return &ImageViewer{
		image,
	}
}

func (viewer *ImageViewer) Display() {
	fmt.Printf("display: %s \n", viewer.data)
}
