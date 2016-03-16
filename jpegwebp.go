package jpegwebp

import (
	"image"
	_ "image/jpeg"
	"io"
	"log"

	"github.com/chai2010/webp"
	"github.com/topicai/candy"
)

func convert() {
	im := candy.WithOpened(candy.TestData("a.jpg"), func(r io.Reader) interface{} {
		im, _, e := image.Decode(r)
		if e != nil {
			log.Panic(e)
		}
		return im
	}).(image.Image)

	candy.WithCreated(candy.TestData("a.webp"), func(w io.Writer) {
		if e := webp.Encode(w, im, nil); e != nil {
			log.Panic(e)
		}
	})
}
