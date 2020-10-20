// Package engine реализует поиск в указанном словаре.
package engine

import (
	"fmt"
	"strings"
)

// Scanner предоставляет возможность получить словарь
type Scanner interface {
	Scan() (map[string]string, error)
}

// Search получает словарь и осуществляет в нем поиск по фразе
func Search(s Scanner, phrase string) ([]string, error) {
	var found []string
	dictionary, err := s.Scan()
	if err != nil {
		return found, err
	}

	for k, v := range dictionary {
		if strings.Contains(k, phrase) || strings.Contains(v, phrase) {
			found = append(found, fmt.Sprintf("%s - '%s'\n", k, v))
		}
	}
	return found, nil
}
