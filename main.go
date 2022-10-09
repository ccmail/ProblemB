package main

import (
	"Huawei/ProblemB/util"
	"fmt"
	"sort"
)

var (
	// BatchNumber 一个批次的数量,一个批次的所有，视为染色体
	BatchNumber int
	// Population 种群
	Population [][]int
)

func main() {

	data := util.ReadCsv()
	sort.Slice(data, func(i, j int) bool {
		if data[i].Width == data[j].Width {
			return data[i].Length < data[j].Length
		}
		return data[i].Width < data[j].Width
	})
	//赋值
	BatchNumber = len(data)
	//初始化种群，创建PopulationNum个个体，染色体为BatchNumber的随机排列
	Population = util.InitPopulation(util.PopulationNum, BatchNumber)
	/*for i := range Population {
		fmt.Println(Population[i])
	}
	*/

	//先迭代300次
	for epoch := 0; epoch < 1; epoch++ {
		//雌雄配队
		male, female := util.Breed(Population)
		newGeneration := make([][]int, 0, util.PopulationNum)

		//在这里由个体的染色体排列得到实际的排列组合情况, 根据实际排列情况进行打分
		//mergeStack()
		//score := util.Score(Population)
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
		mergeStack(data, Population[0])
		//更细世代
		copy(Population, newGeneration)
		//fmt.Println(epoch)
	}

}

//mergeStack 工件列表，以及工件的排列顺序
func mergeStack(items []util.Pair, individual []int) []util.Pair {
	var stack []util.Pair
	//开始拼接，按照染色体拼接
	stackLength, stackWidth := 0.0, 0.0
	//将拼接到的items的id记录下来
	itemIds := make([]int, 0)

	//按照染色体拼接
	for _, v := range individual {
		item := items[v]
		for j := 0; j < item.Count; j++ {
			//当栈的长度将要大于原件的长度时，将其合并程一个栈，栈的长度为所有item的长度，宽度为最大宽度
			if stackLength+item.Length > util.MaxLength {
				//深拷贝，避免反向传播
				copyItemIds := make([]int, len(itemIds))
				copy(copyItemIds, itemIds)
				//再加就大于板材长度了，所哟判定形成了一个单独的栈,加入栈中时处理了下，保证“长>高”
				stack = append(stack, util.Pair{
					util.MaxF(stackLength, stackWidth), util.MinF(stackWidth, stackLength),
					len(itemIds), copyItemIds})
				stackWidth, stackLength, itemIds = 0, 0, itemIds[:0]
			}
			//默认栈是按长度拼接，所以长度是取和，高度取max
			stackLength += item.Length
			//默认工件宽度是小于原料宽度
			stackWidth = util.MaxF(item.Width, stackWidth)
			//将当前item编号加进去
			itemIds = append(itemIds, item.Ids[j])
		}
	}

	fmt.Println(len(stack))
	for i := range stack {
		fmt.Println(stack[i].Length, stack[i].Width)
	}
	//等下合并为stripe的时候按照插入stack的顺序组合，边界按宽度来
	return stack
	/*	for i := 0; i < len(individual); i++ {

		}

		isUsed := make([]int, len(items))
		for i := range items {
			currItem := items[i]
			//当前item还有剩的
			for isUsed[i] < currItem.Count {
				remainCount := currItem.Count - isUsed[i]
				minUseCount := util.MinI(remainCount, int(util.MaxLength/currItem.Length))

				//看能合并出多少个相同的栈
				sameStackCount := remainCount / minUseCount
				//将item合并为栈,并将id组合进去
				for j := 0; j < sameStackCount; j++ {
					stackContainIds := make([]int, 0)
					stackContainIds = append(stackContainIds, currItem.Ids[isUsed[i]:isUsed[i]+j*minUseCount]...)
					stack = append(stack, util.Pair{float64(minUseCount) * currItem.Length, currItem.Width, sameStackCount, stackContainIds})
				}
				isUsed[i] += minUseCount * sameStackCount
			}
		}
		//先仿照, 将栈从大到小排列
		sort.Slice(stack, func(i, j int) bool {
			if stack[i].Length == stack[j].Length {
				return stack[i].Width > stack[j].Length
			}
			return stack[i].Length > stack[j].Length
		})
		return stack*/
}

func mergeStripe([]util.Pair) {

}

/*func mergeStripe(stacks []util.Pair) util.Pair {
	var stripe []util.Pair
	for i := 0; i < len(stacks); i++ {
		currStack := stacks[i]
		for isUsedStackNum := 0; isUsedStackNum < currStack.Count; {
			var stripeWidth float64

			for stripeWidth <= util.MaxWidth {

				if isUsedStackNum > len(stacks) {
					flag := false
					temp := util.Pair{0, 0, 0, []int{}}

					temp.Width, temp.Length = stacks[i].Width, stacks[i].Length
					temp.Ids = append(temp.Ids, stacks[])
				}
			}
		}
	}
}
*/
