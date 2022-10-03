package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Println("Usage : " + os.Args[0] + " file name")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	h := bounds.Max.Y
	w := bounds.Max.X

	bmpi := 0
	bmp := make([]uint8, w*h/8)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			bmpi = x & 7
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()

			if r != 0 || g != 0 || b != 0 {
				bmp[(y*w+x)/8] |= 1 << uint(7-bmpi)
			}
		}
	}

	fmt.Println()

	fileName := path.Base(os.Args[1])
	fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]

	fmt.Printf("#define %v_width %d\n", fileName, w)
	fmt.Printf("#define %v_height %d\n", fileName, h)

	fmt.Printf("const uint8_t %v[] PROGMEM = {\n\t", fileName)
	i := 0
	comma := ", "

	for k, v := range bmp {
		if i > 15 {
			i = 0
			fmt.Printf("\n\t")
		}
		i++
		fmt.Printf("0x%02x%v", v, comma)

		if k == len(bmp)-2 {
			comma = ""
		}
	}

	fmt.Println("\n};")

}
