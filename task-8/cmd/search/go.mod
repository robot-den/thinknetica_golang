module cmd/search

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/crawler/webscnr v1.0.0
replace pkg/crawler/webscnr => ./../../pkg/crawler/webscnr

require pkg/crawler/stubscnr v1.0.0
replace pkg/crawler/stubscnr => ./../../pkg/crawler/stubscnr

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/index/word v1.0.0
replace pkg/index/word => ../../pkg/index/word

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/engine v1.0.0
replace pkg/engine => ../../pkg/engine

require pkg/engine/btree v1.0.0
replace pkg/engine/btree => ../../pkg/engine/btree

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/storage/file v1.0.0
replace pkg/storage/file => ../../pkg/storage/file

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

go 1.15
