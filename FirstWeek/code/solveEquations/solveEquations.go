package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	LEFT      byte = 'L' // 代表在方程等式的左侧
	RIGHT     byte = 'R' // 代表在方程等式的右侧
	NUMERICAL byte = 'N' // 代表运算数
	OPERATOR  byte = 'O' // 代表运算符
	LETTER    byte = 'l' // 代表未知数字母
	UNKNOWN   byte = 'U' // 代表未知字符
	ERROR     byte = 'E' // 代表错误
)

func StingToInt(s interface{}) (int, bool) {
	switch s.(type) {
	case string:
		str, _ := s.(string)
		if len(str) == 1 {
			return int(str[0] - 48), true
		} else {
			i, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Atoi字符串转整型 错误", err)
			}
			return i, true
		}
	default:
		return 0, false
	}
}

func JudgeCharType(c byte) byte {
	// 判断字符c是 数字 还是 字母,返回'N'是数字，'L'是字母，'O'是运算符，'U'是其他字符
	if c >= '0' && c <= '9' {
		return NUMERICAL
	} else if c >= 'a' && c <= 'z' {
		return LETTER
	} else if c == '+' || c == '-' || c == '=' {
		return OPERATOR
	} else {
		return UNKNOWN
	}
}

func JudgeInputStr(inputStr string, unknownNum *byte) (equation string, err error) {
	// 判断输入是否合法
	newStr := make([]byte, 0, len(inputStr)) // 存放规范后的方程字符串
	unknowns := make(map[byte]int, 1)        // 用于统计未知数
	equalSign := map[byte]int{'=': 0}        // 用于统计 '='的个数
	for i := 0; i < len(inputStr); i++ {
		s := inputStr[i]
		switch s {
		case '\n', '+', '-', '=':
			if s == '\n' && i > 0 && JudgeCharType(inputStr[i-1]) == OPERATOR {
				return "", errors.New("+ 、-、= 不能在末尾！")
			} else if s == '=' {
				if i == 0 {
					return "", errors.New(" '=' 不能在开头！")
				} else if i > 0 && JudgeCharType(inputStr[i-1]) == OPERATOR {
					return "", errors.New("+=、-=、==均不合规范")
				}
				equalSign['='] = equalSign['='] + 1
			} else {
				if i > 0 && (JudgeCharType(inputStr[i-1]) == '+' || JudgeCharType(inputStr[i-1]) == '-') {
					return "", errors.New("-+、+-、++、--均不合规范")
				}
			}
			if s != '\n' {
				newStr = append(newStr, s)
			}
			continue
		case ' ':
			continue
		default:
			if JudgeCharType(s) == NUMERICAL {
				if i > 0 && JudgeCharType(inputStr[i-1]) == LETTER { // a2 设置为不合规范，2a规范
					return "", errors.New("未知数与系数书写不规范！")
				}
				newStr = append(newStr, s)
				continue
			} else if JudgeCharType(s) == LETTER {
				if i > 0 && JudgeCharType(inputStr[i-1]) == LETTER { // xx ，aa 不合规范
					return "", errors.New("未知数书写不规范，方程需为一元一次方程！")
				}
				_, ok := unknowns[s]
				if !ok {
					unknowns[s] = 1
				}
				newStr = append(newStr, s)
				continue
			} else if s == '.' {
				return "", errors.New("输入不能包含'.'（或输入不能包含小数）！")
			} else if s >= 'A' && s <= 'Z' {
				return "", errors.New("输入不能包含大写字母！")
			} else if s == '/' {
				return "", errors.New("输入不能包含除号！")
			} else {
				return "", errors.New(fmt.Sprintf("输入不能包含字符：%c", s))
			}
		}
	}
	if len(unknowns) > 1 {
		return "", errors.New("未知数过多！")
	}
	if n, _ := equalSign['=']; n != 1 {
		return "", errors.New(" = 过多！")
	}
	for k, _ := range unknowns {
		*unknownNum = k
		break
	}
	return string(newStr), nil
}

func ConvertNumPosAndNeg(strNumber string, currentState byte, leftOrRight byte, charType byte) int {
	// 转换读取到的等式两边数值的正负
	//currentState 表示数值现在的正负情况
	//leftOrRight表示在等式左边(L)还是右边(R)
	//charType表示该数值是常数还是未知数的系数
	//fmt.Println("LR:", leftOrRight)
	unsignedNumber, _ := StingToInt(strNumber) // 数字型字符串转换为int型

	if leftOrRight == LEFT {
		if charType == LETTER {
			if currentState == '-' {
				return -unsignedNumber
			} else {
				return unsignedNumber
			}
		} else {
			if currentState == '+' {
				return -unsignedNumber
			} else {
				return unsignedNumber
			}
		}

	} else {
		if charType == LETTER {
			if currentState == '+' {
				return -unsignedNumber
			} else {
				return unsignedNumber
			}
		} else {
			if currentState == '-' {
				return -unsignedNumber
			} else {
				return unsignedNumber
			}
		}
	}
}

func ReadInputStr() (equation string, unknownNum byte, err error) {
	// 提醒用户输入并获取输入内容
	fmt.Print("请输入一元一次方程：")
	reader := bufio.NewReader(os.Stdin)
	inputStr, err1 := reader.ReadString('\n')
	if err1 != nil {
		es := fmt.Sprintf("read error:%v", err1)
		return "", ERROR, errors.New(es)
	}
	if len(inputStr) == 0 || inputStr[0] == '\n' {
		return "", ERROR, errors.New("输入为空，请重新输入！")
	}
	// 判断输入是否合法，合法返回规范后的方程字符串
	if equation, err = JudgeInputStr(inputStr, &unknownNum); err != nil {
		return "", ERROR, err
	}
	return equation, unknownNum, nil
}

func ExtractEquationParameters(equationStr string, unknownNum byte) (constants int, coefficient int, err error) {
	// 提取方程参数，提取计算未知数的系数，以及提取常数并计算常数和
	var strLength = len(equationStr)
	var constantsSum = 0                                       // 计算常数和
	var currentOperator byte = '+'                             // 存放当前运算符，除开'='
	var leftOrRight = LEFT                                     // 判断读取进行到方程等式左或右
	var unknownNumAndCoefficient = map[byte]int{unknownNum: 0} // key:未知数 value:未知数系数
	var charType byte                                          // 字符类型
	var num int                                                // 保存读取到并进行相应转化的数值
	for i := 0; i < strLength; i++ {
		charType = JudgeCharType(equationStr[i])
		if charType == NUMERICAL {
			j := i
			for {
				j++
				if j >= strLength {
					s := equationStr[i:j]
					num = ConvertNumPosAndNeg(s, currentOperator, leftOrRight, NUMERICAL)
					constantsSum += num
					i = j
					break
				}
				t := JudgeCharType(equationStr[j])
				if t == NUMERICAL { // 字符为数字，继续读取
					continue
				}
				// 读取到操作符或者未知数的时候停止，进入转换操作
				s := equationStr[i:j] // s : string型，获取的数字字符串切片
				if t == OPERATOR {
					num = ConvertNumPosAndNeg(s, currentOperator, leftOrRight, NUMERICAL) // 根据在等号左右进行数值转换
					currentOperator = equationStr[j]
					i = j
					constantsSum += num
					if equationStr[j] == '=' {
						leftOrRight = RIGHT
					}
					break
				} else if t == LETTER {
					num = ConvertNumPosAndNeg(s, currentOperator, leftOrRight, LETTER) // 根据在等号左右进行数值转换
					unknownNumAndCoefficient[unknownNum] += num
					i = j
					break
				}
			}
		} else if charType == OPERATOR {
			currentOperator = equationStr[i] // 记录当前操作符
			if equationStr[i] == '=' {
				leftOrRight = RIGHT
			}
			continue
		} else if equationStr[i] == unknownNum {
			if i == 0 { // 一开始就读到系数为1的未知数的时候
				unknownNumAndCoefficient[unknownNum] += 1
			} else { // 在其他地方读到系数为1的未知数时
				if leftOrRight == RIGHT {
					if currentOperator == '-' {
						unknownNumAndCoefficient[unknownNum] += 1
					} else {
						unknownNumAndCoefficient[unknownNum] -= 1
					}
				} else {
					if currentOperator == '+' {
						unknownNumAndCoefficient[unknownNum] += 1
					} else {
						unknownNumAndCoefficient[unknownNum] -= 1
					}
				}
			}
			continue
		} else {
			continue
		}
	}
	return constantsSum, unknownNumAndCoefficient[unknownNum], nil
}

func SolveEquations() (equation string, unknownNumber byte, result string) {
	// 解一元一次方程，返回参数列表： 原方程，未知数，解
	var err error
	for {
		equation, unknownNumber, err = ReadInputStr()
		if err != nil {
			fmt.Println(err)
			continue
		}
		constantsSum, coefficient, err1 := ExtractEquationParameters(equation, unknownNumber)
		//fmt.Println("const:", constants)
		//fmt.Println("coefi:", coefficient)
		if err1 != nil {
			fmt.Println(err1)
		}
		if coefficient == 0 {
			if constantsSum == 0 {
				return equation, unknownNumber, "无穷解"
			} else {
				return equation, unknownNumber, "方程无解"
			}
		} else {
			result = fmt.Sprintf("%0.3f", float64(constantsSum)/float64(coefficient))
			return equation, unknownNumber, result
		}
	}
}

func main() {
	equation, unknownNumber, result := SolveEquations()
	fmt.Printf("方程：%v\n解：%c=%v", equation, unknownNumber, result)
}
