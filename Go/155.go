package main

type MinStack struct {
    min   []int
    stack []int
}

func Constructor() MinStack {
    return MinStack{
        min:   []int{},
        stack: []int{},
    }
}

func (this *MinStack) Push(val int) {
    if len(this.stack) != 0 {
        if this.min[len(this.min)-1] > val {
            this.min = append(this.min, val)
        } else {
            this.min = append(this.min, this.min[len(this.min)-1])
        }
    } else {
        this.min = append(this.min, val)
    }
    this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}
