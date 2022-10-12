package main

import (
	"Huawei/ProblemB/util"
	"fmt"
	"math/rand"
	"strings"
)

var (
	// BatchNumber 一个批次的数量,一个批次的所有，视为染色体
	BatchNumber int
	// Population 种群
	Population [][]int
)

func main() {

	files := util.ScanDir()
	for _, file := range files {

		classAndData := util.ReadCsv(file)
		materialCnt := 1
		for material, data := range classAndData {
			/*	item_cnt := 0
				for i := range classAndData {
					item_cnt += classAndData[i].Count
				}*/
			//fmt.Println("item初始数量为", item_cnt)
			//赋值
			BatchNumber = len(data)
			//初始化种群，创建PopulationNum个个体，染色体为BatchNumber的随机排列
			Population = util.InitPopulation(util.PopulationNum, BatchNumber)
			/*for i := range Population {
				fmt.Println(Population[i])
			}
			*/

			//每个材料是一个种群，先迭代300次
			for epoch := 0; epoch < util.EpochNum; epoch++ {

				//首先生成每个个体的分数（概率）
				score := util.Score(Population, data)

				//fmt.Println(len(score), len(Population))
				//fmt.Println(len(Population))

				newGeneration := make([][]int, 0)
				Population, newGeneration = util.PickBest(Population, score)
				//fmt.Println(len(Population))

				//fmt.Println()
				//fmt.Println()
				//fmt.Println()
				//随机挑一个不参加繁殖
				if len(Population)%2 != 0 {
					randIndex := rand.Intn(len(Population))
					newGeneration = append(newGeneration, Population[randIndex])
					Population = append(Population[:randIndex], Population[randIndex+1:]...)
				}
				//雌雄配队
				male, female := util.Breed(Population)

				//在这里由个体的染色体排列得到实际的排列组合情况, 根据实际排列情况进行打分
				//stack存在一个id

				//两两结合， 产子概率0.9，每次产子随机交换父母染色体
				for pairNum := len(male) - 1; pairNum >= 0; pairNum-- {
					maleNum, femaleNum := male[pairNum], female[pairNum]
					son, daughter := util.Cross(Population[maleNum], Population[femaleNum], util.BreedProbability)
					//儿子和女儿进行变异操作，这里有点缺陷，没有繁殖的个体也会参与突变，不过概率在2%
					son = util.Mutation(son, util.MaxMutationNum, util.MutationProbability)
					daughter = util.Mutation(daughter, util.MaxMutationNum, util.MutationProbability)
					//添加到新种群
					newGeneration = append(newGeneration, son, daughter)
				}

				fmt.Println("新世代个体数是", len(newGeneration))

				Population = make([][]int, len(newGeneration))
				//更细世代
				copy(Population, newGeneration)

				currPopulationVal := 0x7ffffff

				currBestPopulation := []util.Pair{}
				//	经过300代繁殖后，挑选当前种群使用钢板最少的
				for _, v := range Population {
					stack := util.MergeStack(data, v)
					stripe := util.MergeStripe(stack)

					if currPopulationVal > len(stripe) {
						currPopulationVal, currBestPopulation = len(stripe), stripe
					}
				}

				fmt.Println("当前种群是", material, "...", "第", epoch+1, "/", util.EpochNum, "次迭代完成，最优的个体使用的钢板数为", len(currBestPopulation))
			}
			currPopulationVal := 0x7ffffff

			currBestPopulation := []util.Pair{}
			//	经过300代繁殖后，挑选当前种群使用钢板最少的
			for _, v := range Population {
				stack := util.MergeStack(data, v)
				stripe := util.MergeStripe(stack)

				if currPopulationVal > len(stripe) {
					currPopulationVal, currBestPopulation = len(stripe), stripe
				}
			}

			fmt.Println()
			fmt.Println()
			fmt.Println()

			split := strings.Split(file, "\\")
			fileName := strings.Split(split[len(split)-1], ".")[0]
			util.OutPutImageAndCsv(currBestPopulation, data, material, fileName, materialCnt)

			materialCnt++
		}

	}
	if util.DataSetSelect == util.CheckA {
		util.ReverseImage(util.OutputPathA)
	} else {
		util.ReverseImage(util.OutputPathB)
	}

}
