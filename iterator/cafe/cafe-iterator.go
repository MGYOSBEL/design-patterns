package cafe

import (
	"fmt"

	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

type CafeMenuIterator struct {
	position int
	elements []*menu.MenuItem
}

func (it *CafeMenuIterator) HasNext() bool {
	return it.position < len(it.elements)
}

func (it *CafeMenuIterator) Next() *menu.MenuItem {
	item := it.elements[it.position]
	it.position++
	return item
}

func NewCafeMenuIterator(items map[string]*menu.MenuItem) iterator.Iterator {
	elements := make([]*menu.MenuItem, 0)
	for _, value := range items {
		elements = append(elements, value)
	}
	fmt.Println("initialized cafe menu iterator values")
	return &CafeMenuIterator{
		position: 0,
		elements: elements,
	}
}
