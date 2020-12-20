module cmd/search

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/crawler/webscnr v1.0.0
replace pkg/crawler/webscnr => ./../../pkg/crawler/webscnr

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ./../../pkg/index/hash

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/engine v1.0.0
replace pkg/engine => ../../pkg/engine

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

require pkg/plugin v1.0.0
replace pkg/plugin => ../../pkg/plugin

require pkg/plugin/netsrv v1.0.0
replace pkg/plugin/netsrv => ../../pkg/plugin/netsrv

require pkg/plugin/webapp v1.0.0
replace pkg/plugin/webapp => ../../pkg/plugin/webapp

go 1.15
