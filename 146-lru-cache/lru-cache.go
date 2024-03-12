
import "container/list"
type LRUCache struct {
    cache *list.List
    capacity int
    value map[int]*list.Element
}

type cacheItem struct {
    key int
    value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{ cache: list.New(), capacity: capacity, value: make(map[int]*list.Element)}
}

func (c *LRUCache) Get(key int) int {
    if _, exists := c.value[key]; !exists {
        return -1
    } else {
        element := c.value[key]
        c.cache.MoveToBack(element)
        return element.Value.(*cacheItem).value
    }
}


func (c *LRUCache) Put(key int, value int)  {
    if element, exists := c.value[key]; exists {
        element.Value.(*cacheItem).value = value
        c.cache.MoveToBack(element)
    } else {
        element = c.cache.PushBack(&cacheItem{ key: key, value: value})
        c.value[key] = element
    }


    if c.cache.Len() > c.capacity {
		// Remove the front element from the cache and the map
		frontElement := c.cache.Front()
		c.cache.Remove(frontElement)
		delete(c.value, frontElement.Value.(*cacheItem).key)
	}

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */