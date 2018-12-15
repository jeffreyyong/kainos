package problem0015

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
	one [][]int
}

func Test_Problem0015(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, -1, -1, 0}},
			ans{[][]int{
				[]int{-1, 0, 1},
			}},
		},

		question{
			para{[]int{-1, 0, 1, 2, 2, 2, 2, -1, -4}},
			ans{[][]int{
				[]int{-4, 2, 2},
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			}},
		},
		question{
			para{[]int{0, 0, 0, 0, 0}},
			ans{[][]int{
				[]int{0, 0, 0},
			}},
		},
		question{
			para{[]int{1, 1, -2}},
			ans{[][]int{
				[]int{-2, 1, 1},
			}},
		},
		question{
			para{[]int{0, 0, 0}},
			ans{[][]int{
				[]int{0, 0, 0},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}
