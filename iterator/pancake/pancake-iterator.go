package pancake

import (
	"github.com/MGYOSBEL/design-patterns/iterator/iterator"
	"github.com/MGYOSBEL/design-patterns/iterator/menu"
)

type PancakeMenuIterator struct {
	position  int
	menuItems []*menu.MenuItem
}

func (it *PancakeMenuIterator) HasNext() bool {
	return (it.position < len(it.menuItems))
}

func (it *PancakeMenuIterator) Next() *menu.MenuItem {
	item := it.menuItems[it.position]
	it.position++
	return item
}

func NewPancakeMenuIterator(items []*menu.MenuItem) iterator.Iterator {
	return &PancakeMenuIterator{
		position:  0,
		menuItems: items,
	}
}
