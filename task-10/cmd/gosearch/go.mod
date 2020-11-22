module cmd/gosearch

go 1.14

replace (
	pkg/api => ../../pkg/api
	pkg/crawler => ../../pkg/crawler
	pkg/crawler/membot => ../../pkg/crawler/membot
	pkg/crawler/spider => ../../pkg/crawler/spider
	pkg/engine => ../../pkg/engine
	pkg/index => ../../pkg/index
	pkg/index/hash => ../../pkg/index/hash
	pkg/storage => ../../pkg/storage
	pkg/storage/memstore => ../../pkg/storage/memstore
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.8.0
	google.golang.org/protobuf v1.25.0 // indirect
	pkg/api v0.0.0-00010101000000-000000000000
	pkg/crawler v0.0.0-00010101000000-000000000000
	pkg/crawler/spider v0.0.0-00010101000000-000000000000
	pkg/engine v0.0.0-00010101000000-000000000000
	pkg/index v0.0.0-00010101000000-000000000000
	pkg/index/hash v0.0.0-00010101000000-000000000000
	pkg/storage v0.0.0-00010101000000-000000000000
	pkg/storage/memstore v0.0.0-00010101000000-000000000000
)
