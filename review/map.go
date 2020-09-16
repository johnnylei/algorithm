package review

import (
	"sync"
	"time"
)

type MyMap struct {
	// TODO: add your own implementation
	// You can use the default golang map
	bucket map[string][]interface{}
	lo sync.RWMutex
	defaultValue string
	defaultValueTime int64
}
func (m *MyMap) Put(key, value string) {
	// TODO: implement the put function to set a specfic key to a specific value.
	// Requirement: need to implement in O(1) time complexity.
	m.lo.Lock()
	defer m.lo.Unlock()
	if _, okay := m.bucket[key]; !okay {
		m.bucket[key] = make([]interface{}, 2)
	}
	m.bucket[key][0] = value
	m.bucket[key][1] = time.Now().UnixNano()
}
func (m *MyMap) Get(key string) (string, bool) {
	// TODO: return the specific map value for the given key. if the key does not exit, return empty with false.
	// Requirement: need to implement in O(1) time complexity.
	m.lo.RLock()
	defer m.lo.RUnlock()
	value, okay := m.bucket[key]
	if !okay {
		if m.defaultValueTime == 0 {
			return "", false
		}

		return m.defaultValue, true
	}

	valueStr, _ := value[0].(string)
	valueTime, _ := value[1].(int64)
	if valueTime < m.defaultValueTime {
		return m.defaultValue, true
	}

	return valueStr, true
}
func (m *MyMap) PutAll(value string) {
	// TODO: set all the existing keys in the map to the given value.
	// Requirement: need to implement in O(1) time complexity.
	m.lo.Lock()
	defer m.lo.Unlock()
	m.defaultValue = value
	m.defaultValueTime = time.Now().UnixNano()
}
