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

func TestRemoveElement(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{[]int{1, 1, 2}, 1, 1},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 1, 7},
	}

	for _, q := range qs {
		param1, param2, answer := q.param1, q.param2, q.answer
		fmt.Printf("~~%v~~~~%v~~\n", param1, param2)

		ast.Equal(answer, RemoveElement(param1, param2), "输入:%v", param1, param2)
	}
}
