package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	param1 []int
	param2 int
	answer int
}

func TestSearchInsert(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3}, 4, 2},
		{[]int{1, 3}, 0, 0},
	}

	for _, q := range qs {
		param1, param2, answer := q.param1, q.param2, q.answer
		fmt.Printf("~~%v~~~~%v~~\n", param1, param2)

		ast.Equal(answer, SearchInsert1(param1, param2), "输入:%v", param1, param2)
	}
}
