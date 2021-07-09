package golru

import (
	"container/list"
	"errors"
)

type LRUCache struct {
	m map[string]*list.Element	// For accessing element in LRU cache's linked list in O(1) time
	sizeLimit int				// Maximum amount of elements that can be stored in LRU Cache
	l list.List					// For adding/removing elements from head or tail in O(1) time
}

type KeyValuePair struct{
	key string			// Key to access value stored in access map
	value interface{}	// The value we will be storing in the access map
}

func createLRU(sizeLimit int) LRUCache {
	return LRUCache{
		m: map[string]*list.Element{},
		sizeLimit : sizeLimit,
		l: list.List{},
	}
}

// GetElem retrieves an element from the LRU Cache
func (lruCache *LRUCache) GetElem(k string) (*list.Element, error) {
	if elem, exists := lruCache.m[k]; !exists {
		return nil, errors.New("element does not exist in cache")
	} else {
		// Element is accessed again therefore move to the front (make most recently used)
		lruCache.l.MoveToFront(elem)
		return elem, nil
	}

}

// SetElem places an element in the LRU Cache, if it is a new element create a node otherwise reassign value
func (lruCache *LRUCache) SetElem(k string, v interface{}) {
	newEntry := KeyValuePair{
		key: k,
		value: v,
	}
	if elem, exists := lruCache.m[k]; !exists {
		// Add new element to the front of the doubly linked list (make most recently used)
		elem = lruCache.l.PushFront(newEntry)
		// Element does not exist in cache, define in map for constant access time O(1)
		lruCache.m[k] = elem

		if lruCache.l.Len() > lruCache.sizeLimit {
			keyToRemove := lruCache.l.Back().Value.(KeyValuePair).key
			// LRU Cache has reached its size limit remove the least recently used element from the access map
			delete(lruCache.m, keyToRemove)
			// LRU Cache has reached its size limit remove the least recently used element from the doubly linked list
			lruCache.l.Remove(lruCache.l.Back())
		}

	} else {
		// Element exists in cache, reassign to the provided value in the access map
		elem.Value = newEntry
		// Relocate element to the front of the doubly linked list (make most recently used)
		lruCache.l.MoveToFront(elem)
	}
}