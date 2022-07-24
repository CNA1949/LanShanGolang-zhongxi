package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadInputOriginalString() (originalString string, offset byte, err error) {
	// 读取原文字符串并进行合法性判断 offset偏移量 originalString 原文字符串
	var n int
label1:
	for {
		fmt.Print("请输入偏移量：")
		fmt.Scanln(&n)
		if n < 0 {
			fmt.Println("偏移量不能为负！")
			continue label1
		}
		offset = byte(n)
		fmt.Printf("请输入原文字符串：")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			return "", 0, errors.New("读取错误")
		}
		originalString = strings.TrimRight(str, "\n")

		// 判断其合法性
		strLength := len(originalString)
		if strLength < 1 || strLength > 50 {
			fmt.Println("输入的原文字符串不能为空！")
			continue label1
		}
		for i := 0; i < strLength-1; i++ {
			if originalString[i] >= 'a' && originalString[i] <= 'z' {
				continue
			} else {
				fmt.Println("输入的原文字符串只能含有小写字母!")
				continue label1
			}
		}
		return originalString, offset, nil
	}
}

func CovertToCipher(originalString string, offset byte) (cipher string, err error) {
	// 根据原文字符串 originalString 和 偏移量 offset 求出密码
	originalStr := []byte(originalString)
	for i := 0; i < len(originalStr); i++ {
		if originalStr[i]+offset > 'z' {
			originalStr[i] = (originalStr[i]+offset)%('z') + byte(96)
		} else {
			originalStr[i] += offset
		}
	}
	cipher = string(originalStr)
	return cipher, nil
}

func main() {
	str, offset, err := ReadInputOriginalString()
	if err != nil {
		fmt.Println(err)
		return
	}
	cipher, _ := CovertToCipher(str, offset)
	fmt.Println("密码：", cipher)
}
