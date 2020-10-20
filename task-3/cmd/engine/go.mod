module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/engine v1.0.0
replace pkg/engine => ../../pkg/engine
// Пришлось добавить это сюда. Вероятно modules не умеют обрабатывать replace в go.mod файлах зависимостей.
require pkg/crawler_stub v1.0.0
replace pkg/crawler_stub => ../../pkg/crawler_stub

go 1.15
