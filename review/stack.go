package review

import "sync"

type MyStack struct {
	// TODO: add your own implementation
	min []int
	lo sync.Mutex
	bucket []int
	length int
}
func (m *MyStack) Push(value int) {
	// Requirement: need to implement in O(1) time complexity.
	m.lo.Lock()
	defer m.lo.Unlock()
	if m.length == 0 {
		m.min = append(m.min, value)
		m.bucket = append(m.bucket, value)
		m.length++
		return
	}

	if value < m.min[m.length-1] {
		m.min = append(m.min, value)
	} else {
		m.min = append(m.min, m.min[m.length - 1])
	}
	m.bucket = append(m.bucket, value)
	m.length++
}
func (m *MyStack) Pop() (int) {
	// Requirement: need to implement in O(1) time complexity.
	m.lo.Lock()
	defer m.lo.Unlock()
	if m.length == 0 {
		panic("empty stack")
	}

	ret := m.bucket[m.length - 1]
	m.bucket = m.bucket[:m.length - 1]
	m.min = m.min[:m.length - 1]
	m.length--
	return ret
}
func (m *MyStack) GetMin() (int) {
	// Requirement: need to implement in O(1) time complexity.
	if m.length == 0 {
		panic("empty stack")
	}

	return m.min[m.length - 1]
}
