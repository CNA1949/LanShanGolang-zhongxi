package main

type AnimalShelf [][]int

//func Constructor() (_ AnimalShelf) { return }

func (a *AnimalShelf) Enqueue(animal []int) {
	if len(animal) < 2 {
		return
	}
	*a = append(*a, animal)
}

func (a *AnimalShelf) DequeueAny() []int {
	var oldest []int
	if len(*a) < 1 {
		return []int{-1, -1}
	}
	oldest = (*a)[0]
	*a = (*a)[1:]
	return oldest
}

func (a *AnimalShelf) DequeueDog() []int {
	if len(*a) < 1 {
		return []int{-1, -1}
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i][1] == 1 {
			return a.DeleteSlice(i)
		}
	}
	return []int{-1, -1}
}

func (a *AnimalShelf) DequeueCat() []int {
	if len(*a) < 1 {
		return []int{-1, -1}
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i][1] == 0 {
			return a.DeleteSlice(i)
		}
	}
	return []int{-1, -1}
}

func (a *AnimalShelf) DeleteSlice(index int) (del []int) {
	// 从切片中删除下标为index的值
	var temp [][]int
	del = (*a)[index]
	temp = (*a)[index+1:]
	*a = (*a)[0:index]
	*a = append(*a, temp...)
	return del
}

