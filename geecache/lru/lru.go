package lru

import "container/list"

/**
* @project: go-gee
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/28/21 3:28 PM
**/

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64                    // 允许使用的最大内存
	nbytes   int64                    // 当前已使用内存
	ll       *list.List               // container包提供的双向链表，LRU实现的数据结构之一
	cache    map[string]*list.Element // 字典，LRU实现的数据结构之二
	// Optional and executed when an entry is purged.
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数，可以为nil
}

// doubly linked list node
type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes.
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}


