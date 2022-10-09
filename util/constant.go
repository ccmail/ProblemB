package util

const (
	// MaxLength 最大长度
	MaxLength float64 = 2440
	// MaxWidth 最大宽度
	MaxWidth float64 = 1220

	/**遗传算法设置**/

	// PopulationNum 种群数量
	PopulationNum int = 400
	// EpochNum 迭代次数
	EpochNum int = 300
	//MaxMutationNum 变异基因数，每次随机变异[1,MaxMutationNum]个,最小值设置应当为1
	MaxMutationNum = 3

	// BreedProbability 交叉概率
	BreedProbability float64 = 0.9
	// MutationProbability 编译概率
	MutationProbability float64 = 0.2
	//	交叉染色体个数
)
