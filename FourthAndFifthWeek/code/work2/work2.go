package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) inToOut() {
	// 依次从inStack末尾去除元素压入outStack中
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1] // 删除inStack末尾元素
	}
}

func (q *MyQueue) Push(x int) {
	// 入队列：将元素 x 推到队列的末尾
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	// 出队列：从队列的开头移除并返回元素
	if len(q.outStack) == 0 {
		q.inToOut() //如果输出栈中没有元素，则从输入栈中取出全部元素
	}
	x := q.outStack[len(q.outStack)-1]          // 输出栈末尾元素即为队列头元素
	q.outStack = q.outStack[:len(q.outStack)-1] //取出后删除
	return x
}

func (q *MyQueue) Peek() int {
	// 返回队列开头的元素
	if len(q.outStack) == 0 {
		q.inToOut() //如果输出栈中没有元素，则从输入栈中取出全部元素
	}
	return q.outStack[len(q.outStack)-1] // 输出栈末尾元素即为队列头元素
}

func (q *MyQueue) Empty() bool {
	// 如果队列为空，返回 true ；否则，返回 false
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func main() {
	obj := Constructor()
	fmt.Println("队列插入元素之前是否为空：", obj.Empty())
	fmt.Println("##### 元素插入队列：")
	for i := 1; i <= 5; i++ {
		fmt.Printf("入队顺序： %d, 入队元素：%d\n", i, i+10)
		obj.Push(i + 10) // 向队列中插入元素
	}
	fmt.Println("对列中元素：", obj.inStack)
	fmt.Println("此时队列是否为空：", obj.Empty())
	fmt.Println("队头元素：", obj.Peek())
	fmt.Println("#####元素出队：")
	for i := 1; i <= 5; i++ {
		fmt.Printf("出队顺序: %d , 出队元素：%v\n", i, obj.Pop())
	}
	fmt.Println("此时队列是否为空：", obj.Empty())
}

