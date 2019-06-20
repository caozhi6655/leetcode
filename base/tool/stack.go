package tool

import "sync"

type Item interface{}

type Stack struct {
	items []Item
	Size  int
	lock  sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.items = []Item{}
	return s
}

func (s *Stack) Push(data interface{}) {
	if s == nil {
		panic("stack is nil")
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, data)
	s.Size++
}

// 出栈
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s == nil {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.Size--
	return item
}

// 最顶元素，不出栈
func (s *Stack) Top() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s == nil {
		return nil
	}

	item := s.items[len(s.items)-1]
	return item
}
