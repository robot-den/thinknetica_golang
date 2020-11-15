module pkg/engine/array

require pkg/model v1.0.0
replace pkg/model => ../../../pkg/model

require pkg/crawler/stubscnr v1.0.0
replace pkg/crawler/stubscnr => ./../../crawler/stubscnr

require pkg/index/word v1.0.0
replace pkg/index/word => ../../../pkg/index/word

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../../pkg/storage/memory

go 1.15