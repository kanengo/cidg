module github.com/kanengo/cidg/example/service2

go 1.25.4

replace (
	github.com/kanengo/cidg => ../../
	github.com/kanengo/cidg/example/service1 => ../service1
)

require github.com/kanengo/cidg v0.0.0
