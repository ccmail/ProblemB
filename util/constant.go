package util

const (
	// MaxLength 最大长度
	MaxLength float64 = 2440
	// MaxWidth 最大宽度
	MaxWidth float64 = 1220

	/**遗传算法设置**/

	// PopulationNum 种群数量
	PopulationNum int = 10000
	// EpochNum 迭代次数
	//EpochNum int = 10000
	EpochNum int = 400
	//MaxMutationNum 变异基因数，每次随机变异[1,MaxMutationNum]个,最小值设置应当为1
	MaxMutationNum = 10

	// BreedProbability 交叉概率
	BreedProbability float64 = 0.8
	// MutationProbability 编译概率
	MutationProbability float64 = 0.3

	// BestNum 轮盘选中个数
	BestNum int = 10

	DataSetSelect = "A"

	CheckA    = "A"
	CheckB    = "B"
	DataPathA = "./data/a"
	DataPathB = "./data/b"

	OutputPathA = "./output/img/a/"
	OutputPathB = "./output/img/b/"
)
