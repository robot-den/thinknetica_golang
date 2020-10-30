// Package btree реализует бинарное дерево для записей типа model.Record
package btree

import (
	"fmt"
	"pkg/model"
)

// BTree представляет собой бинарное дерево, предоставляющее методы по добавлению и поиску элементов
type BTree struct {
	root *node
}

// node представляет собой узел бинарного дерева. Этот тип хранит значение и ссылки на другие узлы
type node struct {
	left  *node
	right *node
	value *model.Record
}

// Add позволяет добавить элемент в бинарное дерево
func (bt *BTree) Add(r *model.Record) {
	newNode := &node{
		value: r,
	}

	if bt.root == nil {
		bt.root = newNode
		return
	}

	parentNode := bt.root
	for {
		if parentNode.value.Id > r.Id {
			if parentNode.left == nil {
				parentNode.left = newNode
				break
			}
			parentNode = parentNode.left
			continue
		}

		if parentNode.right == nil {
			parentNode.right = newNode
			break
		}
		parentNode = parentNode.right
	}
}

// Search осуществляет поиск в бинарном дереве. Второе возвращаемое значение равно false если запись не найдена
func (bt *BTree) Search(id int) (*model.Record, bool) {
	currentNode := bt.root

	for {
		if currentNode == nil {
			return &model.Record{}, false
		}

		if currentNode.value.Id == id {
			return currentNode.value, true
		}

		if currentNode.value.Id > id {
			currentNode = currentNode.left
			continue
		}

		currentNode = currentNode.right
	}
}

// String позволяет получить простое строковое представление бинарного дерева
func (bt *BTree) String() string {
	elems := []int{}
	bt.root.collect(&elems)
	return fmt.Sprint(elems)
}

// collect выполняет рекурсивный обход дерева и собирает Id элементов в массив
func (n *node) collect(s *[]int) {
	if n == nil {
		return
	}

	*s = append(*s, n.value.Id)

	n.left.collect(s)
	n.right.collect(s)
}
