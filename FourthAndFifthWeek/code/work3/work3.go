package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	var operand = make([]int, 0, len(s))
	var operator = make([]uint8, 0, len(s))
	operator = append(operator, '+')
	var temp = ""
	// 处理字符串
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp += string(s[i])
		} else {
			switch s[i] {
			case '+', '-', '*', '/':
				v, _ := strconv.Atoi(temp)
				temp = ""

				if len(operator)-1 >= 0 && operator[len(operator)-1] == '*' {
					operand[len(operand)-1] *= v
					operator[len(operator)-1] = s[i]
					continue
				} else if len(operator)-1 >= 0 && operator[len(operator)-1] == '/' {
					operand[len(operand)-1] /= v
					operator[len(operator)-1] = s[i]
					continue
				}
				operator = append(operator, s[i])
				operand = append(operand, v)
			}
		}
		if i == len(s)-1 {
			v, _ := strconv.Atoi(temp)
			if len(operator)-1 >= 0 {
				switch operator[len(operator)-1] {
				case '*':
					operand[len(operand)-1] *= v
					if len(operator) > 1 {
						operator = operator[:len(operator)-1]
					}
				case '/':
					operand[len(operand)-1] /= v
					if len(operator) > 1 {
						operator = operator[:len(operator)-1]
					}
				case '+', '-':
					operand = append(operand, v)
				}
			}
			break
		}
	}

	// 运算
	var result = 0
	for i := 0; i < len(operand); i++ {
		if operator[i] == '-' {
			result -= operand[i]
		} else {
			result += operand[i]
		}
	}
	return result
}

func main() {
	fmt.Printf("算式：\"0\", 结果：%v\n", calculate("0"))
	fmt.Printf("算式：\"1+1\", 结果：%v\n", calculate("1+1"))
	fmt.Printf("算式：\"3+2*2\", 结果：%v\n", calculate("3+2*2"))
	fmt.Printf("算式：\" 3/2 \", 结果：%v\n", calculate(" 3/2 "))
	fmt.Printf("算式：\" 3+5 / 2 \", 结果：%v\n", calculate(" 3+5 / 2 "))
	fmt.Printf("算式：\"3*6/2\", 结果：%v\n", calculate("3*6/2"))
}
