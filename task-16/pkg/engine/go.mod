module pkg/engine

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

go 1.15
