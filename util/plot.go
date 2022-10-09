package util

import (
	"fmt"
	"github.com/fogleman/gg"
	"strconv"
)

var (
	dc    *gg.Context
	white = "white"
	black = "black"
	red   = "red"
)

func OutPutImage(stripe, item []Pair) {
	dc = gg.NewContext(2800, 1400)

	dc.SetHexColor("#000000")
	dc.Clear()

	//原料板子, api的width是板子的长,height是板子的宽
	currStartX, currStartY := 0.0, 0.0
	//drawRectangleBackBound(currStartX, currStartY, 2440/4, 1220/4, white)

	totalImage := 0
	imageCnt := 0
	for _, v := range stripe {

		if imageCnt > 16 {

			dc.SavePNG("out" + strconv.Itoa(totalImage) + ".png")
			totalImage++
			imageCnt = 0
			currStartX, currStartY = 0.0, 0.0
			dc = gg.NewContext(2800, 1400)
			dc.SetHexColor("#000000")
			dc.Clear()
		}

		drawMaxWidth, drawMaxLength, recordWidth := 0.0, 0.0, 0.0
		drawRectangleBackBound(currStartX, currStartY, 2440/4, 1220/4, white)
		colorIdx := 0
		for _, itemId := range v.Ids {
			if drawMaxLength+item[itemId].Length > MaxLength {
				recordWidth += drawMaxWidth
				drawMaxLength = 0
				drawMaxWidth = 0
			}
			drawRectangleLine(currStartX+(drawMaxLength/4), currStartY+recordWidth/4,
				item[itemId].Length/4, item[itemId].Width/4, colorIdx)
			colorIdx++
			drawMaxLength += item[itemId].Length
			drawMaxWidth = MaxF(drawMaxWidth, item[itemId].Width)

		}
		imageCnt++
		currStartX = float64((imageCnt % 4) * 700)
		currStartY = float64((imageCnt / 4) * 360)
		fmt.Println(v.Length, v.Width)
	}

	dc.StrokePreserve()
	//dc.SetRGB(0, 0, 0)
	//dc.Fill()

}

func drawRectangleBackBound(x float64, y float64, w float64, h float64, color string) {
	dc.DrawRectangle(x, y, w, h)
	if color == white {
		dc.SetHexColor("#ffffff")
	} else if color == black {
		dc.SetHexColor("#000000")
	}
	dc.Fill()
}

func drawRectangleLine(x, y, w, h float64, color int) {
	color %= 4
	dc.DrawRectangle(x, y, w, h)
	if color == 0 {
		dc.SetHexColor("#800000")
	} else if color == 1 {
		dc.SetHexColor("#7B68EE")
	} else if color == 2 {
		dc.SetHexColor("#008B8B")
	} else if color == 3 {
		dc.SetHexColor("#FF8C00")
	}
	dc.Fill()
	//dc.SetHexColor("#ffffff")
	//dc.SetLineWidth(5)
	//dc.StrokePreserve()
	//dc.Stroke()
}
