package model

import "sync"

type autoIncr struct {
	sync.Mutex
	id int
}

func (a *autoIncr) GetID() (id int) {
	a.Lock()
	defer a.Unlock()
	id = a.id
	a.id++
	return
}
