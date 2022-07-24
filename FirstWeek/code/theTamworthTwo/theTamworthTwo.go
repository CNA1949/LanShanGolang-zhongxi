package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	EAST  byte = 'E' // 向东走
	WEST  byte = 'W' // 向西走
	SOUTH byte = 'S' // 向南走
	NORTH byte = 'N' // 向北走
)

type Coordinate struct { // 人或牛的坐标
	X int
	Y int
}

func JudgeMap(analogousMap [][]byte) (fCoordinate Coordinate, cCoordinate Coordinate, err error) {
	// 判断输入是否合法
	// 是否只存在 '.'、'*'、'C'、'F'四种字符
	// 保证10行，每行10个字符
	// 'C'和'F' 均只存在一个
	var cCount int // 统计 'C' 的字符个数
	var fCount int // 统计 'F' 的字符个数
	iLength := len(analogousMap)
	if iLength < 10 {
		return fCoordinate, cCoordinate, errors.New("地图行大小不正确！")
	}
	for i := 0; i < len(analogousMap); i++ {
		jLength := len(analogousMap[i])
		if jLength < 10 {
			return fCoordinate, cCoordinate, errors.New("地图列大小不正确！")
		}
		for j := 0; j < jLength; j++ {
			c := analogousMap[i][j]
			switch c {
			case '.', '*':
				continue
			case 'C':
				cCoordinate = Coordinate{X: i, Y: j}
				cCount++
				continue
			case 'F':
				fCoordinate = Coordinate{X: i, Y: j}
				fCount++
				continue
			default:
				return fCoordinate, cCoordinate, errors.New("地图符号不规范！")
			}
		}
	}
	if cCount != 1 || fCount != 1 {
		return fCoordinate, cCoordinate, errors.New("F 和 C 只能出现一次！")
	}
	//if fCoordinate.X == cCoordinate.X && fCoordinate.Y == cCoordinate.Y {
	//	return fCoordinate, cCoordinate, errors.New("F 和 C 起点位置不能相同！")
	//}// 在终端输入地图 ，F 和 C 不可能起点在同一位置
	return fCoordinate, cCoordinate, nil
}

func ReadInputMap() (analogousMap [][]byte, farmerCoordinate Coordinate, cowCoordinate Coordinate, err error) {
	var maps = make([][]byte, 10)

	for {
		fmt.Println("请输入模拟地图：")
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < 10; i++ {
			inputLine, err1 := reader.ReadBytes('\n')
			if err1 != nil {
				es := fmt.Sprintf("read error:%v", err1)
				return analogousMap, Coordinate{}, Coordinate{}, errors.New(es)
			}
			inputLine = []byte(strings.TrimRight(string(inputLine), "\n"))
			maps[i] = inputLine
		}
		var err2 error
		farmerCoordinate, cowCoordinate, err2 = JudgeMap(maps)
		if err2 != nil {
			fmt.Println()
			fmt.Println("输入不合法:", err2)
			continue
		}
		break
	}
	return maps, farmerCoordinate, cowCoordinate, nil
}

func NextCoordinate(analogousMap [][]byte, currentCoordinate *Coordinate, currentDirection *byte) {
	// 计算人或牛下一个移动的坐标
	switch *currentDirection {
	case EAST:
		if (*currentCoordinate).Y+1 >= 10 || analogousMap[(*currentCoordinate).X][(*currentCoordinate).Y+1] == '*' {
			*currentDirection = SOUTH
		} else {
			(*currentCoordinate).Y += 1
		}
	case WEST:
		if (*currentCoordinate).Y-1 < 0 || analogousMap[(*currentCoordinate).X][(*currentCoordinate).Y-1] == '*' {
			*currentDirection = NORTH
		} else {
			(*currentCoordinate).Y -= 1
		}
	case NORTH:
		if (*currentCoordinate).X-1 < 0 || analogousMap[(*currentCoordinate).X-1][(*currentCoordinate).Y] == '*' {
			*currentDirection = EAST
		} else {
			(*currentCoordinate).X -= 1
		}
	case SOUTH:
		if (*currentCoordinate).X+1 >= 10 || analogousMap[(*currentCoordinate).X+1][(*currentCoordinate).Y] == '*' {
			*currentDirection = WEST
		} else {
			(*currentCoordinate).X += 1
		}
	}
}

func Tracking(analogousMap [][]byte, currentFCoordinate Coordinate, currentCCoordinate Coordinate) (isCatch bool, costTime int) {
	iMax := 4 * 10 * 10 * 4 * 10 * 10 // 超出 iMax 次追逐，追到的可能性为0
	fDirection := NORTH               // 起始朝向为北方
	cDirection := NORTH               // 起始朝向为北方
	costTime = 0                      // 人追逐到牛所花费的时间（分钟）
	for i := 1; i < iMax; i++ {
		NextCoordinate(analogousMap, &currentCCoordinate, &cDirection) // 计算牛的下一个位置
		NextCoordinate(analogousMap, &currentFCoordinate, &fDirection) // 计算人的下一个位置
		costTime++
		if currentFCoordinate.X == currentCCoordinate.X && currentFCoordinate.Y == currentCCoordinate.Y {
			return true, costTime
		}
	}
	return false, 0
}

func main() {
	maps, farmerCoordinate, cowCoordinate, err := ReadInputMap()
	if err != nil {
		fmt.Println("error:", err)
	}
	isCatch, cost := Tracking(maps, farmerCoordinate, cowCoordinate)
	if isCatch {
		fmt.Printf("能够抓住，耗时：%v 分钟", cost)
	} else {
		fmt.Println("抓不住牛！")
	}
}
