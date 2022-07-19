package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	listItem, ok := c.items[key]
	if !ok {
		listItem = c.queue.PushFront(cacheItem{
			key:   key,
			value: value,
		})
		if c.queue.Len() > c.capacity {
			leastRecentlyUsedItem := c.queue.Back()
			c.queue.Remove(leastRecentlyUsedItem)
			v, ok := leastRecentlyUsedItem.Value.(cacheItem)
			if ok {
				delete(c.items, v.key)
			}
		}
		c.items[key] = listItem
		return false
	}
	listItem.Value = cacheItem{
		key:   key,
		value: value,
	}
	c.queue.MoveToFront(listItem)
	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	listItem, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(listItem)
		v, ok := listItem.Value.(cacheItem)
		if ok {
			return v.value, true
		}
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
