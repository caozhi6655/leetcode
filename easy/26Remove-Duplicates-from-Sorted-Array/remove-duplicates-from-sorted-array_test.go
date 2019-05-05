package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	param  []int
	answer int
}

func TestRemoveDuplicates(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, q := range qs {
		param, answer := q.param, q.answer
		fmt.Printf("~~%v~~\n", param)

		ast.Equal(answer, RemoveDuplicates(param), "输入:%v", param)
	}
}
