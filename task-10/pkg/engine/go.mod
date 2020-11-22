module pkg/engine

go 1.15

replace (
	pkg/crawler => ../crawler
	pkg/index => ../index
	pkg/storage => ../storage
	pkg/index/hash => ../index/hash
)

require (
	pkg/crawler v0.0.0-00010101000000-000000000000
	pkg/index v0.0.0-00010101000000-000000000000
	pkg/storage v0.0.0-00010101000000-000000000000
)
