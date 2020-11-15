module pkg/index/word

require pkg/model v1.0.0
replace pkg/model => ../../../pkg/model

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../../pkg/storage/memory

require pkg/crawler/stubscnr v1.0.0
replace pkg/crawler/stubscnr => ./../../crawler/stubscnr

go 1.15
