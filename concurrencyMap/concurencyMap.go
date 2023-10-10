package concurrencyMap

import "sync"

type ConcurrencyMap[K comparable, V any] struct {
	Map map[K]V
	sync.Mutex
}

func NewConcurrencyMap[K comparable, V any]() *ConcurrencyMap[K, V] {
	return &ConcurrencyMap[K, V]{
		Map: make(map[K]V),
	}
}

// Set прежде, чем записывать значение value по ключу k, - лочим на чтение и запись
func (cm *ConcurrencyMap[K, V]) Set(key K, value V) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map[key] = value
}
