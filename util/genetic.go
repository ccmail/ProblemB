package util

import (
	"math/rand"
)

// Cross 交叉两个个体的基因
func Cross(x, y []int, p float64) ([]int, []int) {
	currProbability := RandFloat64()
	//以随机概率进行交配
	if currProbability > p {
		return x, y
	}
	n := len(x)
	//开始挑选对那些染色体进行交叉，这里采用部分匹配交叉，保证全排列
	l := rand.Intn(n)
	r := rand.Intn(n-l) + l
	for l == r {
		l = rand.Intn(n)
		r = rand.Intn(n-l) + l
	}
	xMp := make(map[int]int)
	yMp := make(map[int]int)

	for i := l; i < r; i++ {
		x[i], y[i] = y[i], x[i]
	}

	for i := l; i < r; i++ {
		xMp[x[i]] = y[i]
		yMp[y[i]] = x[i]
	}

	unique(x, l, r, n, xMp)
	unique(y, l, r, n, yMp)

	//fmt.Println(x, y)
	//交叉序了，返回去
	return x, y
}

//保证交叉的唯一性
func unique(x []int, l, r, n int, mp map[int]int) {
	for i := 0; i < l; i++ {
		pre := x[i]
		for curr, ok := mp[pre]; ok; curr, ok = mp[curr] {
			pre = curr
		}
		x[i] = pre
	}
	for i := r; i < n; i++ {
		pre := x[i]
		for curr, ok := mp[pre]; ok; curr, ok = mp[curr] {
			pre = curr
		}
		x[i] = pre
	}

}

// Breed 挑选两个个体进行繁殖
func Breed(x [][]int) ([]int, []int) {
	//种群中个体数
	n := len(x)
	male := make([]int, n>>1)
	female := make([]int, n>>1)

	for i := range male {
		female[i], male[i] = (n>>1)+i, i
	}

	female = ShuffleInt(female)
	male = ShuffleInt(male)

	//fmt.Println(male, female)
	return male, female
}

// Mutation 使用交换突变，随机选择两个或者多个基因交换
func Mutation(x []int, mutationNum int, p float64) []int {
	currMutationNum := MaxI(rand.Intn(mutationNum), 1)
	currProbability := RandFloat64()
	//以随机概率进行变异，大于p（init0.1）时不变异
	if currProbability > p {
		return x
	}
	n := len(x)

	for i := 0; i < currMutationNum; i++ {
		//随机选择两个点变异
		l, r := rand.Intn(n), rand.Intn(n)
		for l == r {
			l, r = rand.Intn(n), rand.Intn(n)
		}
		x[l], x[r] = x[r], x[l]
	}
	return x
}

// InitPopulation 初始化种群
func InitPopulation(popNum, chromosomeNum int) [][]int {
	//创建种群数组
	ans := make([][]int, popNum)

	for i := range ans {
		//Population[i] = make([]float64, BatchNumber)
		temp := make([]int, chromosomeNum)
		for j := range temp {
			temp[j] = j
		}
		//随机打乱每一个个体的染色体
		//ans[i] = util.Shuffle(temp)
		ans[i] = ShuffleInt(temp)
	}
	return ans
}

/*func Score(population [][]int, stripes []Pair, currIdx int) []int {
	res := make([]int, PopulationNum)
	genesNum := len(population[0])
	//minScore := 0x7fffffff
	for i := 0; i < PopulationNum; i++ {
		//遍历种族中的所有个体
		currIndividual := population[i]
		//得分和stripe的总长度
		score, totalLen := 0, 0.0
		//开始遍历基因
		for j := 0; j < genesNum; j++ {
			idx := currIndividual[j] + currIdx
			//判断这么多基因排列再一起需要多少块板子,推测是得分越高优先级越低
			stripesLength := stripes[idx].Length
			if totalLen+stripesLength > MaxLength {
				score++
				totalLen = stripesLength
			} else {
				totalLen += stripesLength
			}
		}
		res[i] = score
	}
	return res
}
*/

// Score 采用轮盘赌方法, 分数上,使用材料数量最少的权重最大,
func Score(population [][]int, items []Pair) []float64 {
	materialNum := make([]int, len(population))
	sum := 0
	for i, pop := range population {
		stack := MergeStack(items, pop)
		stripe := MergeStripe(stack)
		materialNum[i] = len(stripe)
		sum += materialNum[i]
	}
	score := make([]float64, len(population))
	for i, v := range materialNum {
		score[i] = (float64(v) / float64(sum)) * float64(sum-v)
	}
	//累加概率
	probabilitySum := 0.0
	for i, v := range score {
		probabilitySum += v
		score[i] = probabilitySum
	}
	return score
}
