module pkg/index

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

go 1.15
