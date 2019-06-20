package main

func main() {

}

type LinkList struct {
	Data int
	Next *LinkList
}

func New(data int) *LinkList {
	return &LinkList{Data: data, Next: nil}
}

//在第i个位置插入data
func (head *LinkList) Insert(i int, data int) {
	l := head
	j := 0
	for l != nil && j < i {
		l = l.Next
		j++
	}

	if l == nil {
		panic("i is too long")
	}

	p := &LinkList{Data: data}
	p.Next = l.Next
	l.Next = p
}
