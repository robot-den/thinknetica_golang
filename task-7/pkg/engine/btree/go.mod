module pkg/engine/btree

require pkg/model v1.0.0
replace pkg/model => ../../../pkg/model

require pkg/btree v1.0.0
replace pkg/btree => ../../../pkg/btree

require pkg/crawler/stub v1.0.0
replace pkg/crawler/stub => ../../../pkg/crawler/stub

require pkg/index/word v1.0.0
replace pkg/index/word => ../../../pkg/index/word

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../../pkg/storage/memory

go 1.15
