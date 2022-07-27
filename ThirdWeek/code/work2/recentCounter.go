package main

import "fmt"

type RecentCounter struct {
	requests []int
}

//func Constructor() (_ RecentCounter) { return }

func (r *RecentCounter) Ping(t int) int {
	r.requests = append(r.requests, t)
	for r.requests[0] < t-3000 {
		r.requests = r.requests[1:]
	}
	return len(r.requests)
}

func test(ping [][]int) {
	var r RecentCounter
	var results []int
	fmt.Println("输入：", ping)
	for _, v := range ping {
		if len(v) < 1 {
			results = append(results, 0)
			continue
		}
		results = append(results, r.Ping(v[0]))
	}
	fmt.Printf("输出：%v\n\n", results)
}

func main() {
	test([][]int{{}, {1}, {100}, {3001}, {3002}})
	test([][]int{{}, {1}, {100}, {2001}, {5002}, {6000}})
	test([][]int{{}, {1}, {100}, {3001}, {4002}, {6003}, {9004}})
}
