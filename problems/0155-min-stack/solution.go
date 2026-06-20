// Package minstack solves LeetCode 155. Min Stack.
// https://leetcode.com/problems/min-stack/
package minstack

type Node struct {
	value    int
	minValue int
}

type MinStack struct {
	store []Node
}

func Constructor() MinStack {
	return MinStack{store: make([]Node, 0)}
}

func (this *MinStack) Push(value int) {
	node := Node{value: value, minValue: value}
	stackLen := len(this.store)

	if stackLen != 0 && value > this.store[stackLen-1].minValue {
		node.minValue = this.store[stackLen-1].minValue
	}

	this.store = append(this.store, node)
}

func (this *MinStack) Pop() {
	this.store = this.store[:len(this.store)-1]
}

func (this *MinStack) Top() int {
	return this.store[len(this.store)-1].value
}

func (this *MinStack) GetMin() int {
	return this.store[len(this.store)-1].minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
