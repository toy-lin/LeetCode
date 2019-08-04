package min_stack

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

import "sync"

type MinStack struct {
	Top_    *Node
	minList []int
}

var Pool = &sync.Pool{}

type Node struct {
	Value int
	Next  *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Top_: nil}
}

func (this *MinStack) Push(x int) {
	if this == nil {
		return
	}
	var n *Node
	p := Pool.Get()
	if p == nil {
		n = &Node{}
	} else {
		n = p.(*Node)
	}
	n.Value = x
	n.Next = this.Top_
	this.Top_ = n
	if len(this.minList) == 0 || this.minList[len(this.minList)-1] > x {
		this.minList = append(this.minList, x)
	} else {
		this.minList = append(this.minList, this.minList[len(this.minList)-1])
	}
}

func (this *MinStack) Pop() {
	if this == nil || this.Top_ == nil {
		return
	}
	Pool.Put(this.Top_)
	this.Top_ = this.Top_.Next
	this.minList = this.minList[:len(this.minList)-1]
}

func (this *MinStack) Top() int {
	if this == nil || this.Top_ == nil {
		return 0
	}
	return this.Top_.Value
}

func (this *MinStack) GetMin() int {
	if this == nil || len(this.minList) == 0 {
		return 0
	}
	return this.minList[len(this.minList)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// -------------------以下是slice实现------------------------

//type MinStack struct {
//	stack   []int
//	minList []int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{}
//}
//
//func (this *MinStack) Push(x int) {
//	if this == nil {
//		return
//	}
//	this.stack = append(this.stack, x)
//	if len(this.minList) == 0 || this.minList[len(this.minList)-1] > x {
//		this.minList = append(this.minList, x)
//	} else {
//		this.minList = append(this.minList, this.minList[len(this.minList)-1])
//	}
//}
//
//func (this *MinStack) Pop() {
//	if this == nil {
//		return
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//	this.minList = this.minList[:len(this.minList)-1]
//}
//
//func (this *MinStack) Top() int {
//	if this == nil {
//		return 0
//	}
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if this == nil || len(this.minList) == 0 {
//		return 0
//	}
//	return this.minList[len(this.minList)-1]
//}
