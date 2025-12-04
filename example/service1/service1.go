package service1

import "github.com/kanengo/cidg/example/pkg/util"

func Service1() {
	util.Add(1, 2)
}

func Service1Mul(a, b int) int {
	return util.Mul(a, b)
}
