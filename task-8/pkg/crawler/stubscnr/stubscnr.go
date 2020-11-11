// Package stubscnr реализует заглушку сканера содержимого веб-сайтов.
package stubscnr

// StubScnr имитирует Crawler
type StubScnr struct{}

// New создает новый экземпляр типа Stub
func New() *StubScnr {
	stub := StubScnr{}
	return &stub
}

// Scan возвращает статический словарь из трех пар
func (c *StubScnr) Scan() (data map[string]string, err error) {
	data = map[string]string{
		"http://one.com/":   "A little copying is better than a little dependency.",
		"http://two.com/":   "Design the architecture, name the components, document the details.",
		"http://three.com/": "Burn the heretic, kill the mutant, purge the unclean.",
	}

	return data, nil
}
