package problem0053

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one int
}

func Test_Problem0053(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans{6},
		},

		question{
			para{[]int{-2}},
			ans{-2},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxSubArray(p.one), "输入:%v", p)
	}
}
