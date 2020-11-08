// Package stub реализует заглушку сканера содержимого веб-сайтов.
package stub

// Stub имитирует Crawler
type Stub struct{}

// New создает новый экземпляр типа Stub
func New() *Stub {
	stub := Stub{}
	return &stub
}

// Scan возвращает статический словарь из трех пар
func (c *Stub) Scan() (data map[string]string, err error) {
	data = map[string]string{
		"http://one.com/":   "A little copying is better than a little dependency.",
		"http://two.com/":   "Design the architecture, name the components, document the details.",
		"http://three.com/": "Burn the heretic, kill the mutant, purge the unclean.",
	}

	return data, nil
}
