package datastruct

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

func (_self *Node) String() string {
	return fmt.Sprintf("%v", _self.value)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock sync.RWMutex
}

func (_self *Graph) AddNode(node *Node) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.nodes = append(_self.nodes, node)
}

func (_self *Graph) AddEdge(u, v *Node)  {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	if _self.edges == nil {
		_self.edges = make(map[Node][]*Node)
	}

	_self.edges[*u] = append(_self.edges[*u], v)
	_self.edges[*v] = append(_self.edges[*v], v)
}

func (_self *Graph) String()  {
	_self.lock.RLock()
	defer _self.lock.RUnlock()
	str := ""
	for _, iNode := range _self.nodes {
		str += fmt.Sprintf("%s->", iNode.String())
		nexts := _self.edges[*iNode]
		if len(nexts) == 0 {
			continue
		}

		for _, next := range nexts {
			str = fmt.Sprintf("%s %s", str, next.String())
		}
		str += "\n"
	}
	fmt.Println(str)
}