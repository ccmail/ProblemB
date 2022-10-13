package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type pair struct {
	material    string
	utilization float64
}

func calcUtilization(filePath string, startIdx int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("\n读取文件时发生了错误, 错误信息如下： \n", err)
	}
	defer file.Close()

	in := csv.NewReader(file)
	//跳过标题
	in.Read()
	var preMaterial string
	var preOriginalIndex int

	var material string
	var originalIndex int

	currUtilization := make([]float64, 0)
	ans := make([]pair, 0)
	currTotalArea := 0.0
	for {
		//line, err := in.Read()
		line, err := in.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("读取csv行时发生了错误，错误信息如下： \n", err)
		}
		//材质名
		material = line[startIdx+0]

		//原片序号
		originalIndex, _ = strconv.Atoi(line[startIdx+1])

		if material != preMaterial {

			if material == "ZYF-0215S" {
				//fmt.Println("在这里打断")
			}

			latestUtilization := 0.0
			for _, v := range currUtilization {
				latestUtilization += v
			}
			latestUtilization /= float64(len(currUtilization))

			if !math.IsNaN(latestUtilization) {
				ans = append(ans, pair{preMaterial, latestUtilization})
				preMaterial = material
				currUtilization = currUtilization[:0]
				currTotalArea = 0
				preOriginalIndex = originalIndex
			}
		}

		if originalIndex != preOriginalIndex {
			currUtilization = append(currUtilization, currTotalArea/(2440.0*1220.0))
			currTotalArea, preOriginalIndex = 0, originalIndex
		}
		itemLength, _ := strconv.ParseFloat(line[startIdx+5], 64)
		itemWidth, _ := strconv.ParseFloat(line[startIdx+6], 64)
		currTotalArea += itemWidth * itemLength
	}

	if currTotalArea != 0 {
		currUtilization = append(currUtilization, currTotalArea/(2440.0*1220.0))
	}

	if len(currUtilization) != 0 {
		latestUtilization := 0.0
		for _, v := range currUtilization {
			latestUtilization += v
		}
		latestUtilization /= float64(len(currUtilization))
		if !math.IsNaN(latestUtilization) {
			ans = append(ans, pair{preMaterial, latestUtilization})
		}
	}

	ans = ans[1:]
	fmt.Println(len(ans))
	latest := 0.0
	for _, v := range ans {
		latest += v.utilization
		//v.utilization
		//fmt.Println(v.material, "...", v.utilization)
	}
	fmt.Println("===================================================")
	fmt.Println(latest / float64(len(ans)))
}
