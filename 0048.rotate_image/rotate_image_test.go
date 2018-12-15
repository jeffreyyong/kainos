package problem0048

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
	one [][]int
}

type ans struct {
	one [][]int
}

func Test_Problem0048(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			}},
			ans{[][]int{
				[]int{3, 1},
				[]int{4, 2},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 1, 1, 1},
				[]int{2, 2, 2, 2},
				[]int{3, 3, 3, 3},
				[]int{4, 4, 4, 4},
			}},
			ans{[][]int{
				[]int{4, 3, 2, 1},
				[]int{4, 3, 2, 1},
				[]int{4, 3, 2, 1},
				[]int{4, 3, 2, 1},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		rotate(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
