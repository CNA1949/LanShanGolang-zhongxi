package main

import (
	"fmt"
	"sort"
	"strconv"
)

var result = make(map[string]int)

func splitNaturalNumber(preNum int, n int, str string) {
	// 拆分自然数 n
	if n == 1 {
		str = str + "1"
		result[str] = 1
		return
	} else if n == 2 {
		str = str + "1+1"
		result[str] = 1
		return
	}
	newStr := str
	for i := 1; i <= n/2; i++ {
		for j := n - 1; j >= n-i; j-- {
			if i+j == n {
				if i < preNum || j < preNum {
					continue
				}
				originalStr := str
				str = str + strconv.Itoa(i) + "+" + strconv.Itoa(j)
				result[str] = 1
				preNum = i
				str = originalStr + strconv.Itoa(i) + "+"
				k1 := j / 2
				k2 := j - k1
				if k1 >= i && k2 >= i {
					splitNaturalNumber(preNum, j, str)
				}
			}
		}
		str = newStr
	}
}

func main() {
	var n int
	fmt.Print("任意输入一个大于1的自然数：")
	fmt.Scanln(&n)
	splitNaturalNumber(1, n, "")
	fmt.Println("---------拆分结果---------")
	s := make([]string, 0, 10)
	for k, _ := range result {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println("-------------------")
}
