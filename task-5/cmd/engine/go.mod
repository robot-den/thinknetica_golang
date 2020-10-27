module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/stub v1.0.0
replace pkg/stub => ../../pkg/stub

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

go 1.15
