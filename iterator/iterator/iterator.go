package iterator

import "github.com/MGYOSBEL/design-patterns/iterator/menu"

type Iterator interface {
	HasNext() bool
	Next() *menu.MenuItem
}
