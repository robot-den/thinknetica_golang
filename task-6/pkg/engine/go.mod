module pkg/engine

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

go 1.15
