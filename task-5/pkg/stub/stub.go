// Package stub предоставляет заглушки сервисов для тестирования
package stub

// Scanner имитирует получателя методов который соответствует интерфейсу Scanner
type Scanner struct{}

// NewScanner создает новый экземпляр типа Scanner
func NewScanner() *Scanner {
	stub := Scanner{}
	return &stub
}

// Scan возвращает статический словарь из трех пар
func (c *Scanner) Scan() (data map[string]string, err error) {
	data = map[string]string{
		"http://one.com/":   "A little copying is better than a little dependency.",
		"http://two.com/":   "Design the architecture, name the components, document the details.",
		"http://three.com/": "Burn the heretic, kill the mutant, purge the unclean.",
	}

	return data, nil
}
