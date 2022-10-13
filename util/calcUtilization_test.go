package util

import (
	"fmt"
	"testing"
)

func TestOutPutImage(t *testing.T) {
	//str, _ := os.Getwd()
	//fmt.Println(str)
	calcUtilization(Problem1stCsv, 0)
	fmt.Println("===========================================================")
	calcUtilization(Problem2ndCsv, 1)

}
