package datastruct

import "sync"

type StackNode struct {
	value interface{}
}

type Stack struct {
	nodes []*StackNode
	lock  sync.RWMutex
}

func (_self *Stack) Push(node *StackNode) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.nodes = append(_self.nodes, node)
}

func (_self *Stack) Pop() *StackNode {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	length := len(_self.nodes)
	if length == 0 {
		return nil
	}

	ret := _self.nodes[length - 1]
	_self.nodes = _self.nodes[:length - 1]
	return ret
}

func (_self *Stack) IsEmpty() bool  {
	if len(_self.nodes) == 0 {
		return true
	}

	return false
}
