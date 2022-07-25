package main

import (
	"fmt"
)

type Point struct { // 存放轨迹点
	X int
	Y int
}

var (
	row, col  int          // 棋盘行列数
	result    [][]int      // 存放最短路径
	mark      [][]int      // 用来存放每个点是否已经走过
	route     []Point      // 用来临时存放马走的轨迹
	direction = [8][2]int{ // 表示马在每个点能够走的方向
		{-2, 1}, {-1, 2}, {1, 2}, {2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}
)

func check(x, y int) bool {
	// 判断马走的下一个点是否已经走过或者出界
	if x > row || y > col || x < 1 || y < 1 || mark[x-1][y-1] != 0 {
		return true
	}
	return false
}

func popSlice(slice []Point) []Point {
	// 删除切片最后一个元素
	slice = slice[0 : len(slice)-1]
	return slice
}

func outputRoute() {
	// 用于输出马儿到某个点的轨迹
	rl := len(route)
	for i := 0; i < rl-1; i++ {
		p := route[i]
		fmt.Printf("(%v,%v)->", p.X, p.Y)
	}
	fmt.Printf("(%v,%v)  len=%v\n", route[rl-1].X, route[rl-1].Y, rl-1)
}

func formatOutputResult() {
	// 格式化输出最终马到达各个点最少要走几步的结果
	fmt.Println("--------------------")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%v\t", result[i][j])
		}
		fmt.Println()
	}
}

func countHorseDistance(currentX, currentY, endX, endY, dep int) {
	// 计算马走“日”字的最短路径
	var nextX, nextY int
	for i := 0; i < len(direction); i++ {
		nextX = currentX + direction[i][0]
		nextY = currentY + direction[i][1]
		// 判断新坐标是否出界或者是否走过
		if !check(nextX, nextY) {
			route = append(route, Point{nextX, nextY})
			mark[nextX-1][nextY-1] = 1
			if dep == row*col || (nextX == endX && nextY == endY) {
				if nextX == endX && nextY == endY {
					//outputRoute() // 输出当前轨迹
					l := len(route)
					if result[endX-1][endY-1] == -1 || l-1 < result[endX-1][endY-1] {
						result[endX-1][endY-1] = l - 1
					}
				}
			} else {
				countHorseDistance(nextX, nextY, endX, endY, dep)
			}
			mark[nextX-1][nextY-1] = 0 // 回溯时恢复当前点的标记
			route = popSlice(route)    // 回溯时将当前点从轨迹切片中移除
		}
	}
}

func initChessboard(n, m, startX, startY int) {
	// 初始化棋盘和马儿的位置
	row = n
	col = m
	result = make([][]int, n)
	mark = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		mark[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i][j] = -1
		}
	}
	mark[startX-1][startY-1] = 1 // 标记初始位置
	route = make([]Point, 0, row+col)
	route = append(route, Point{startX, startY})
}

func run(n, m, sx, sy int) {
	initChessboard(n, m, sx, sy)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if sx == i && sy == j {
				result[i-1][j-1] = 0 //
				continue
			}
			countHorseDistance(sx, sy, i, j, 1)
		}
	}
	formatOutputResult()
}

func main() {
	var n, m, sx, sy int
	fmt.Print("输入n m x y，用空格隔开：")
	fmt.Scanln(&n, &m, &sx, &sy)
	run(n, m, sx, sy)
}
