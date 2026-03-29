type MinStack struct {
	min int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64,
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
	} else {
		this.stack = append(this.stack, val - this.min)
		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if x < 0 {
		this.min = this.min-x
	}
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack)-1]
	if top > 0 {
		return top + this.min
	}
	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
