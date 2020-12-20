// package webapp реализует подключаемый плагин, который запускает web-сервер для доступа к хранилищу
package webapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"pkg/model"
)

// index представляет собой контракт индексатора, из которого плагин webapp прочитает содержимое
type index interface {
	ReadAll() map[string][]int
}

// storage представляет собой контракт хранилища, из которого плагин webapp прочитает содержимое
type storage interface {
	ReadAll() []model.Document
}

// WebApp представляет собой плагин
type WebApp struct {
	index   index
	storage storage
	address string
}

// New позволяет создать новый объект плагина с заданными настройками
func New(ind index, str storage, address string) *WebApp {
	wa := WebApp{
		index:   ind,
		storage: str,
		address: address,
	}

	return &wa
}

// Run запускает службу для обслуживания запросов и позволяет плагину соответствовать интерфейсу plugin.Service
func (wa *WebApp) Run() {
	router := wa.endpoints()
	err := http.ListenAndServe(wa.address, router)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleRoot обслуживает корневой запрос
func (wa *WebApp) HandleRoot(w http.ResponseWriter, r *http.Request) {
	rootPage := `<h1>Links:</h1><p><a href="/index">Index</a></p><p><a href="/docs">Docs</a></p>`
	_, err := w.Write([]byte(rootPage))
	if err != nil {
		fmt.Println(err)
	}
}

// HandleDocsIndex обслуживает запрос индекса документов
func (wa *WebApp) HandleDocsIndex(w http.ResponseWriter, r *http.Request) {
	index := wa.index.ReadAll()
	templateStr := `<h1>Index</h1>{{ range $key, $value := . }}<li><strong>{{ $key }}</strong>: {{ $value }}</li>{{ end }}`
	tmpl := template.Must(template.New("DocsIndex").Parse(templateStr))
	err := tmpl.Execute(w, index)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleDocs обслуживает запрос списка документов
func (wa *WebApp) HandleDocs(w http.ResponseWriter, r *http.Request) {
	index := wa.storage.ReadAll()
	templateStr := `<h1>Docs</h1>{{ range . }}<li><strong>{{ .ID }}</strong>: {{ .Title }} ({{ .URL }})</li>{{ end }}`
	tmpl := template.Must(template.New("Docs").Parse(templateStr))
	err := tmpl.Execute(w, index)
	if err != nil {
		fmt.Println(err)
	}
}

// endpoints регистрирует обработчики и возвращает роутер
func (wa *WebApp) endpoints() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/index", wa.HandleDocsIndex).Methods(http.MethodGet)
	router.HandleFunc("/docs", wa.HandleDocs).Methods(http.MethodGet)
	router.HandleFunc("/", wa.HandleRoot).Methods(http.MethodGet)

	return router
}
