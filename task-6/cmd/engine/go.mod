module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/engine v1.0.0
replace pkg/engine => ../../pkg/engine

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree

require pkg/storage/file v1.0.0
replace pkg/storage/file => ../../pkg/storage/file

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

go 1.15
