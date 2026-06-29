type MinStack struct {
	Stack []int
	MinStack []int
}	

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.MinStack) == 0 || this.MinStack[len(this.MinStack)-1] >= val {
		this.MinStack = append(this.MinStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}

	n := len(this.Stack)-1
	val := this.Stack[n]
	if val == this.MinStack[len(this.MinStack)-1] {
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
	this.Stack = this.Stack[:n]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}

	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Stack) == 0 {
		return 0
	}

	return this.MinStack[len(this.MinStack)-1]
}
