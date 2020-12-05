module cmd/netsrv

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../pkg/crawler

require pkg/crawler/webscnr v1.0.0
replace pkg/crawler/webscnr => ./../../../pkg/crawler/webscnr

require pkg/index v1.0.0
replace pkg/index => ../../../pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ./../../../pkg/index/hash

require pkg/model v1.0.0
replace pkg/model => ../../../pkg/model

require pkg/engine v1.0.0
replace pkg/engine => ../../../pkg/engine

require pkg/storage v1.0.0
replace pkg/storage => ../../../pkg/storage

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../../pkg/storage/memory

go 1.15