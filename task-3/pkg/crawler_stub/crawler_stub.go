// Package crawlerStub позволяет получить автоматически сгенерированный словарь
package crawlerStub

// Stub имитирует получателя методов
type crawlerStub struct{}

// New создает новый экземпляр типа crawlerStub
func New() crawlerStub {
	return crawlerStub{}
}

// Scan возвращает статический словарь из трех пар
func (c crawlerStub) Scan() (data map[string]string, err error) {
	data = map[string]string{
		"One":   "A little copying is better than a little dependency.",
		"Two":   "Design the architecture, name the components, document the details.",
		"Three": "Burn the heretic, kill the mutant, purge the unclean.",
	}

	return data, nil
}
