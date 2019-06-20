package easy

type MinStack struct {
	items []int
	Size  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Size: 0}
}

func (this *MinStack) Push(x int) {
	if this.Size == 0 {
		this.items = append(this.items, x)
		this.items = append(this.items, x)
	} else {
		min := this.GetMin()
		this.items = append(this.items, x)
		if x > min {
			this.items = append(this.items, min)
		} else {
			this.items = append(this.items, x)
		}
	}
	this.Size++
}

func (this *MinStack) Pop() {
	this.items = this.items[0 : len(this.items)-2]
	this.Size--
}

func (this *MinStack) Top() int {
	if this.Size == 0 {
		return 0
	} else {
		return this.items[len(this.items)-2]
	}
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1]
}
