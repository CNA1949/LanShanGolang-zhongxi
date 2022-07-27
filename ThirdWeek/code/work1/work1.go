package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	var kTime int
	for i := 0; i < len(tickets); i++ {
		if i < k {
			if tickets[i] <= tickets[k] {
				kTime += tickets[i]
			} else {
				kTime += tickets[k]
			}
		} else if i == k {
			kTime += tickets[k]
		} else {
			if tickets[i] >= tickets[k] {
				kTime += tickets[k] - 1
			} else {
				kTime += tickets[i]
			}
		}
	}
	return kTime
}

func main() {
	var t1 = []int{2, 3, 2}
	var t2 = []int{5, 1, 1, 1}
	var t3 = []int{8, 5, 6, 4, 2, 6, 1}
	fmt.Printf("输入：tickets = %v, k = %v\n输出：%v\n\n", t1, 2, timeRequiredToBuy(t1, 2))
	fmt.Printf("输入：tickets = %v, k = %v\n输出：%v\n\n", t2, 0, timeRequiredToBuy(t2, 0))
	fmt.Printf("输入：tickets = %v, k = %v\n输出：%v\n\n", t3, 3, timeRequiredToBuy(t3, 3))
}
