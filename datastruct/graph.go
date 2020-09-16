package datastruct

import (
	"fmt"
	"sync"
)

func InitDefault() *Graph {
	A := &GraphNode{
		value:"A",
	}
	B := &GraphNode{
		value:"B",
	}
	C := &GraphNode{
		value:"C",
	}
	D := &GraphNode{
		value:"D",
	}
	E := &GraphNode{
		value:"E",
	}
	F := &GraphNode{
		value:"F",
	}
	graph := &Graph{}
	graph.AddNode(A)
	graph.AddNode(B)
	graph.AddNode(C)
	graph.AddNode(D)
	graph.AddNode(E)
	graph.AddNode(F)
	graph.AddEdge(A, B)
	graph.AddEdge(A, C)
	graph.AddEdge(D, B)
	graph.AddEdge(D, C)
	graph.AddEdge(D, E)
	graph.AddEdge(D, F)
	graph.AddEdge(E, C)
	return graph
}

type GraphNode struct {
	value string
}

func (_self *GraphNode) String() string {
	return fmt.Sprintf("%v", _self.value)
}

type Graph struct {
	nodes []*GraphNode
	edges map[GraphNode][]*GraphNode
	lock sync.RWMutex
}

func (_self *Graph) AddNode(node *GraphNode) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.nodes = append(_self.nodes, node)
}

func (_self *Graph) AddEdge(u, v *GraphNode)  {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	if _self.edges == nil {
		_self.edges = make(map[GraphNode][]*GraphNode)
	}

	_self.edges[*u] = append(_self.edges[*u], v)
	_self.edges[*v] = append(_self.edges[*v], u)
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

// 队列, 广度优先
func (_self *Graph) BFS()  {
	_self.lock.RLock()
	defer _self.lock.RUnlock()
	que := Que{
		nodes:make([]*QueNode, 0),
	}
	head := _self.nodes[0]
	queNode := &QueNode{
		value:head,
	}
	que.Push(queNode)
	visited := make(map[*GraphNode]bool)
	visited[head] = true
	ret := make([]*GraphNode, 0)
	for {
		if que.IsEmpty() {
			break
		}

		queNode = que.Pop()
		node, _ := queNode.value.(*GraphNode)
		ret = append(ret, node)
		for _, next := range _self.edges[*node] {
			if visited[next] {
				continue
			}

			que.Push(&QueNode{
				value:next,
			})
			visited[next] = true
		}
	}
	for _, item := range ret {
		fmt.Println(item.String())
	}
}

func (_self *Graph) Find(nodeValue string) *GraphNode {
	for _, node := range _self.nodes {
		if node.value == nodeValue {
			return node
		}
	}

	return nil
}

// 栈，深度优先
func (_self *Graph) DFS(headValue string) ([]*GraphNode, map[*GraphNode]*GraphNode) {
	_self.lock.RLock()
	defer _self.lock.RUnlock()
	stack := &Stack{
		nodes:make([]*StackNode, 0),
	}
	head := _self.Find(headValue)
	if head == nil {
		return nil, nil
	}

	stack.Push(&StackNode{
		value:head,
	})
	visited := make(map[*GraphNode]bool)
	visited[head] = true
	ret := make([]*GraphNode, 0)
	parent := make(map[*GraphNode]*GraphNode)
	parent[head] = nil
	for {
		if stack.IsEmpty() {
			break
		}

		stackNode := stack.Pop()
		node, _ := stackNode.value.(*GraphNode)
		ret = append(ret, node)
		for _, next := range _self.edges[*node] {
			if visited[next] {
				continue
			}

			parent[next] = node
			stack.Push(&StackNode{
				value:next,
			})
			visited[next] = true
		}
	}
	return ret, parent
}

// 任意两个节点间，最小路径
func (_self *Graph) MinPath(source, destination string) []*GraphNode {
	destinationNode := _self.Find(destination)
	if destinationNode == nil {
		return nil
	}

	_, parent := _self.DFS(source)
	if parent == nil {
		return nil
	}

	if source == destination {
		return []*GraphNode{
			_self.Find(source),
		}
	}

	path := []*GraphNode{}
	node := destinationNode
	for {
		path = append(path, node)
		if node.value == source {
			break
		}
		node = parent[node]
	}
	return path
}

