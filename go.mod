module github.com/kanengo/cidg

go 1.24.9

replace github.com/kanengo/cidg/example/service1 => ./example/service1

require (
	github.com/goccy/go-yaml v1.19.0
	github.com/kanengo/cidg/example/service1 v0.0.0
	golang.org/x/tools v0.39.0
)

require (
	golang.org/x/mod v0.30.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
)
