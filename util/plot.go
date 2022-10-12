package util

import (
	"encoding/csv"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	dc    *gg.Context
	white = "white"
	black = "black"
	red   = "red"
)

func OutPutImageAndCsv(stripe, item []Pair, material, fileName string, batchCnt int) {

	var csvName string
	if DataSetSelect == CheckA {
		csvName = "./output/cut_program.csv"
	} else {
		csvName = "./output/sum_order.csv"
	}

	file, fileIsExist := os.Open(csvName)
	//如果文件不存在
	if fileIsExist != nil {
		file, _ = os.Create(csvName)
		//中文乱码解决
		file.WriteString("\xEF\xBB\xBF")
	} else {
		file, _ = os.OpenFile(csvName, os.O_APPEND|os.O_RDWR, 0666)
	}
	defer file.Close()
	csvFile := csv.NewWriter(file)
	if fileIsExist != nil {
		if DataSetSelect == CheckA {
			csvFile.Write([]string{"原片材质", "原片序号", "原片id", "产品x坐标", "产品y坐标", "产品x方向长度", "产品y方向长度"})
		} else {
			csvFile.Write([]string{"批次序号", "原片材质", "原片序号", "原片id", "产品x坐标", "产品y坐标", "产品x方向长度", "产品y方向长度"})
		}

	}
	csvFile.Flush()
	defer csvFile.Flush()
	dc = gg.NewContext(2800, 1400)

	dc.SetHexColor("#000000")
	dc.Clear()

	//原料板子, api的width是板子的长,height是板子的宽
	currStartX, currStartY := 0.0, 0.0
	//drawRectangleBackBound(currStartX, currStartY, 2440/4, 1220/4, white)

	totalImage := 0
	imageCnt := 0

	materialCnt := 1
	for _, v := range stripe {

		if imageCnt > 16 {
			var rootPath string
			if DataSetSelect == CheckA {
				rootPath = OutputPathA
			} else {
				rootPath = OutputPathB
			}
			path := rootPath + fileName + material + "out" + strconv.Itoa(totalImage) + ".png"
			dc.SavePNG(path)
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

			writeToCsvLine := []string{
				material, strconv.Itoa(materialCnt), strconv.Itoa(item[itemId].originalIds[0]),
				strconv.FormatFloat(currStartX+drawMaxLength, 'f', 1, 64),
				strconv.FormatFloat(currStartY+recordWidth, 'f', 1, 64),
				strconv.FormatFloat(item[itemId].Length, 'f', 1, 64),
				strconv.FormatFloat(item[itemId].Width, 'f', 1, 64)}

			if DataSetSelect != CheckA {
				writeToCsvLine = append([]string{strconv.Itoa(materialCnt)}, writeToCsvLine...)
			}
			csvFile.Write(writeToCsvLine)

			drawRectangleLine(currStartX+(drawMaxLength/4), currStartY+recordWidth/4,
				item[itemId].Length/4, item[itemId].Width/4, colorIdx)
			colorIdx++
			drawMaxLength += item[itemId].Length
			drawMaxWidth = MaxF(drawMaxWidth, item[itemId].Width)

		}
		materialCnt++
		imageCnt++
		currStartX = float64((imageCnt % 4) * 700)
		currStartY = float64((imageCnt / 4) * 360)
		//fmt.Println(v.Length, v.Width)
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

}

func ReverseImage(path string) {
	//root := "./output/img"
	images := make([]string, 0)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		images = append(images, path)
		return nil
	})
	images = images[1:]
	if err != nil {
		panic(err)
	}
	//for _, file := range images {
	//	fmt.Println(file)
	//}

	for idx, imageName := range images {
		file, err := os.Open(imageName)
		if err != nil {
			log.Fatal("\n读取图片时发生了错误, 错误信息如下： \n", err)
		}
		defer file.Close()
		img, _, _ := image.Decode(file)
		bounds := img.Bounds()
		newRGBImage := image.NewRGBA(bounds)
		x, y := bounds.Dx(), bounds.Dy()
		for i := 0; i < x; i++ {

			for l, r := 0, y-1; l < r; l, r = l+1, r-1 {
				left := img.At(i, l)
				l_r, l_g, l_b, l_a := left.RGBA()
				l_r, l_g, l_b = l_r>>8, l_g>>8, l_b>>8
				right := img.At(i, r)
				r_r, r_g, r_b, r_a := right.RGBA()
				r_r, r_g, r_b = r_r>>8, r_g>>8, r_b>>8

				newRGBImage.SetRGBA(i, l, color.RGBA{uint8(r_r), uint8(r_g), uint8(r_b), uint8(r_a)})
				newRGBImage.SetRGBA(i, r, color.RGBA{uint8(l_r), uint8(l_g), uint8(l_b), uint8(l_a)})

			}
			fmt.Println("图片调整中....", "第", idx, "张...", i, "/", x)
		}

		outFile, _ := os.Create(imageName)
		defer outFile.Close()
		png.Encode(outFile, newRGBImage)
		fmt.Println("调整完了")
	}

}
