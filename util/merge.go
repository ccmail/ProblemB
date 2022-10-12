package util

//MergeStack 工件列表，以及工件的排列顺序
func MergeStack(items []Pair, individual []int) []Pair {
	//sort.Slice(items, func(i, j int) bool {
	//	return items[i].Width*items[i].Length > items[j].Width*items[j].Length
	//})
	var stack []Pair
	//开始拼接，按照染色体拼接
	stackLength, stackWidth := 0.0, 0.0
	//将拼接到的items的id记录下来
	itemIds := make([]int, 0)
	itemOriginalIds := make([]int, 0)

	//按照染色体拼接
	for _, v := range individual {
		item := items[v]
		for j := 0; j < item.Count; j++ {
			//当栈的长度将要大于原件的长度时，将其合并程一个栈，栈的长度为所有item的长度，宽度为最大宽度
			if stackLength+item.Length > MaxLength {
				//深拷贝，避免反向传播
				copyItemIds := make([]int, len(itemIds))
				copy(copyItemIds, itemIds)

				copyItemOriginalIds := make([]int, len(itemOriginalIds))
				copy(copyItemOriginalIds, itemOriginalIds)
				//再加就大于板材长度了，所哟判定形成了一个单独的栈,加入栈中时处理了下，保证“长>高”
				stack = append(stack, Pair{
					MaxF(stackLength, stackWidth), MinF(stackWidth, stackLength),
					len(itemIds), copyItemIds, copyItemOriginalIds})
				stackWidth, stackLength, itemIds, itemOriginalIds = 0, 0, itemIds[:0], itemOriginalIds[:0]
			}
			//默认栈是按长度拼接，所以长度是取和，高度取max
			stackLength += item.Length
			//默认工件宽度是小于原料宽度
			stackWidth = MaxF(item.Width, stackWidth)
			//将当前item编号加进去
			itemIds = append(itemIds, item.Ids[j])
			itemOriginalIds = append(itemOriginalIds, item.originalIds[j])
		}
	}
	//收尾
	if len(itemIds) != 0 {
		//深拷贝，避免反向传播
		copyItemIds := make([]int, len(itemIds))
		copy(copyItemIds, itemIds)

		copyItemOriginalIds := make([]int, len(itemOriginalIds))
		copy(copyItemOriginalIds, itemOriginalIds)
		//再加就大于板材长度了，所哟判定形成了一个单独的栈,加入栈中时处理了下，保证“长>高”
		stack = append(stack, Pair{
			MaxF(stackLength, stackWidth), MinF(stackWidth, stackLength),
			len(itemIds), copyItemIds, copyItemOriginalIds})
		/*stack = append(stack, Pair{
		stackLength, stackWidth,
		len(itemIds), copyItemIds})
		//*/stackWidth, stackLength, itemIds, itemOriginalIds = 0, 0, itemIds[:0], itemOriginalIds[:0]
	}

	cnt := 0
	//fmt.Println("stack的个数是", len(stack))
	for _, v := range stack {
		cnt += len(v.Ids)
		//fmt.Println(v.Length, v.Width)
	}
	//fmt.Println("stack里的item有", cnt)
	//等下合并为stripe的时候按照插入stack的顺序组合，边界按宽度来
	return stack
}

//MergeStripe 这里有两个选择，按照顺序组成stripe，或者再使用一次遗传算法，再使用遗传算法复杂度应该会爆炸
func MergeStripe(stacks []Pair) []Pair {
	//sort.Slice(stacks, func(i, j int) bool {
	//	return stacks[i].Width*stacks[i].Length > stacks[j].Width*stacks[j].Length
	//})
	var stripe []Pair
	stripeLength, stripeWidth := 0.0, 0.0
	stackIds := make([]int, 0)
	stackOriginalIds := make([]int, 0)
	for _, stack := range stacks {
		if stripeWidth+stack.Width > MaxWidth {
			copyStacksIds := make([]int, len(stackIds))
			copy(copyStacksIds, stackIds)

			copyStackOriginalIds := make([]int, len(stackOriginalIds))
			copy(copyStackOriginalIds, stackOriginalIds)
			stripe = append(
				stripe,
				Pair{stripeLength, stripeWidth,
					len(stackIds), copyStacksIds, copyStackOriginalIds})
			stripeLength, stripeWidth, stackIds, stackOriginalIds = 0, 0, stackIds[:0], stackOriginalIds[:0]
		}
		stripeLength = MaxF(stripeLength, stack.Length)
		stripeWidth += stack.Width
		stackIds = append(stackIds, stack.Ids...)
		stackOriginalIds = append(stackOriginalIds, stack.originalIds...)
	}
	//收尾
	if len(stackIds) != 0 {
		copyStacksIds := make([]int, len(stackIds))
		copy(copyStacksIds, stackIds)

		copyStackOriginalIds := make([]int, len(stackOriginalIds))
		copy(copyStackOriginalIds, stackOriginalIds)
		stripe = append(
			stripe,
			Pair{stripeLength, stripeWidth,
				len(stackIds), copyStacksIds, copyStackOriginalIds})
		stripeLength, stripeWidth, stackIds, stackOriginalIds = 0, 0, stackIds[:0], stackOriginalIds[:0]

	}

	//fmt.Println("stripe的个数是", len(stripe))
	cnt := 0
	for _, v := range stripe {
		cnt += len(v.Ids)
	}
	//fmt.Println("stripe中的item个数是", cnt)
	return stripe
}
