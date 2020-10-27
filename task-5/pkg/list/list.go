// Package list реализует двусвязный список вместе с базовыми операциями над ним.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New - конструктор нового списка.
func New() *List {
	return &List{}
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e *Elem) *Elem {
	if l.root == nil {
		e.prev = e
		e.next = e
		l.root = e

		return e
	}

	// Сохраняем старое значение root для удобства
	oldRoot := l.root
	// Новый root: next указывает на старый root, prev указывает куда ранее указывал старый root
	e.next = oldRoot
	e.prev = oldRoot.prev
	l.root = e
	// Старый root: если раньше старый root next указывал на себя, то начнет указывать на новый root, если next указывал
	// на конец списка то этот конец начнет указывать на новый root
	// prev просто указывает на новый root
	oldRoot.prev.next = l.root
	oldRoot.prev = l.root

	return e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	root := l.root
	if root == nil {
		return ""
	}

	s := fmt.Sprint(root.Val)
	current := root.next
	for current != root {
		s += fmt.Sprintf(" %v", current.Val)
		current = current.next
	}

	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	if l.root == nil {
		return l
	}

	if l.root == l.root.next {
		l.root = nil
		return l
	}
	// Сохраняем новый root для удобства
	newRoot := l.root.next
	// Новый root prev начинает ссылаться на конец списка (куда раньше указывал старый root)
	newRoot.prev = l.root.prev
	// Конец списка начинает указывать на новый root вместо старого
	newRoot.prev.next = newRoot
	// Корень списка теперь новый root, на старый больше никто не ссылается
	l.root = newRoot

	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	if l.root == nil || l.root == l.root.next {
		return l
	}

	newRoot := l.root.prev

	elem := l.root
	for {
		elem.prev, elem.next = elem.next, elem.prev
		elem = elem.prev // следующий элемент, берем prev посколько обменяли ссылки местами

		if elem == l.root {
			break
		}
	}

	l.root = newRoot

	return l
}

// DebugPrint печатает список в формате удобном для отладки (со ссылками)
func (l *List) DebugPrint() {
	fmt.Println("------")
	root := l.root

	fmt.Printf("prev: %v, curr: %v, next: %v \n", root.prev.Val, root.Val, root.next.Val)
	current := root.next
	for current != l.root {
		fmt.Printf("prev: %v, curr: %v, next: %v \n", current.prev.Val, current.Val, current.next.Val)
		current = current.next
	}
	fmt.Println("------")
}
