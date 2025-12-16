package service2

import (
	"github.com/kanengo/cidg/example/pkg/util"
	"github.com/kanengo/cidg/example/service1"
)

func Service2(a, b int) int {
	service1.Service1Mul(a, b)
	return util.Add(a, b)
}

// Service2Mul
func Service2Mul(a, b int) int {
	return util.Mul(a, b)
}
