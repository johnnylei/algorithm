package datastruct

import "sync"

type QueNode struct {
	value interface{}
}

type Que struct {
	nodes []*QueNode
	lock sync.RWMutex
}

func (_self *Que) Push(node *QueNode) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.nodes = append(_self.nodes, node)
}

func (_self *Que) Pop() *QueNode {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	if len(_self.nodes) == 0 {
		return nil
	}

	ret := _self.nodes[0]
	_self.nodes = _self.nodes[1:]
	return ret
}

func (_self *Que) IsEmpty() bool  {
	if len(_self.nodes) == 0 {
		return true
	}

	return false
}
