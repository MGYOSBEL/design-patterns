package diner

import (
	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

type DinerMenuIterator struct {
	position  int
	menuItems [MAX_ITEMS]*menu.MenuItem
}

func (it *DinerMenuIterator) HasNext() bool {
	return (it.position < MAX_ITEMS) && (it.menuItems[it.position] != nil)
}

func (it *DinerMenuIterator) Next() *menu.MenuItem {
	item := it.menuItems[it.position]
	it.position++
	return item
}

func NewDinerMenuIterator(items [MAX_ITEMS]*menu.MenuItem) iterator.Iterator {
	return &DinerMenuIterator{
		position:  0,
		menuItems: items,
	}
}
