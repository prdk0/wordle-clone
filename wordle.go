package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	title  string  = "Wordle"
	width  int     = 435
	height int     = 600
	rows   int     = 6
	cols   int     = 5
	dpi    float64 = 72
)

var (
	fontSize        float64 = 24
	mplusNormalFont font.Face
	bkg             = color.White
	lightGrey       = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey            = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow          = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green           = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fontColor       = color.Black
	edge            = false
	alphabet        = "qwertyuiopasdfghjklzxcvbnm"
	grid            [cols * rows]string
	dict            []string
	checkCol        [cols * rows]int
	loc             int = 0
	won                 = false
	answer          string
)

type Game struct {
	rune []rune
}

func main() {
	g := &Game{}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(width, height)
}

