package util

import (
	"encoding/csv"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
//files []string
//root = "./data/a"
)

func ScanDir() (files []string) {
	var root string
	if DataSetSelect == CheckA {
		root = DataPathA
	} else {
		root = DataPathB
	}

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	//for _, file := range files {
	//fmt.Println(file)
	//}
	files = files[1:]
	return files
}

func ReadCsv(files string) (ans map[string][]Pair) {
	file, err := os.Open(files)
	idMap := make(map[int]int)
	if err != nil {
		log.Fatal("\n读取文件时发生了错误, 错误信息如下： \n", err)
	}
	defer file.Close()
	materialMap := make(map[string][]Pair)
	in := csv.NewReader(file)
	//跳过标题
	in.Read()
	id := 0
	for {
		//line, err := in.Read()
		line, err := in.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("读取csv行时发生了错误，错误信息如下： \n", err)
		}
		originalId, _ := strconv.Atoi(line[0])
		length, _ := strconv.ParseFloat(line[3], 64)
		width, _ := strconv.ParseFloat(line[4], 64)
		material := line[1]
		//调整item的长宽，保证长比宽大
		length = MaxF(length, width)
		width = MinF(length, width)
		idMap[id] = originalId

		materialMap[material] = append(materialMap[material],
			Pair{length, width, 1, []int{len(materialMap[material])}, []int{originalId}})
		id++
	}

	return materialMap
}

/*func ReadCsv(files string) (ans []Pair, idMap map[int]int, material string) {
	//files := ScanDir()
	//for _, filePath := range files {
	//	file, err := os.Open(filePath)
	file, err := os.Open(files)
	idMap = make(map[int]int)
	//var material string
	if err != nil {
		log.Fatal("\n读取文件时发生了错误, 错误信息如下： \n", err)
	}
	defer file.Close()

	//res := make(map[[2]float64][]int)
	//csv中的扫描器
	in := csv.NewReader(file)
	//跳过标题
	in.Read()
	id := 0
	for {
		//line, err := in.Read()
		line, err := in.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("读取csv行时发生了错误，错误信息如下： \n", err)
		}
		originalId, _ := strconv.Atoi(line[0])
		length, _ := strconv.ParseFloat(line[3], 64)
		width, _ := strconv.ParseFloat(line[4], 64)
		material = line[1]
		//调整item的长宽，保证长比宽大
		length = MaxF(length, width)
		width = MinF(length, width)

		idMap[id] = originalId
		ans = append(ans, Pair{
			length, width, 1, []int{id}, []int{originalId},
		})
		id++
		//res[[2]float64{length, width}] = append(res[[2]float64{length, width}], id)
	}
	return ans, idMap, material
}
*/
