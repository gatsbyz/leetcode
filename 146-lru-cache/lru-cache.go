type Node struct {
    prev, next *Node
    key, value int
}

type LRUCache struct {
    cacheMap          map[int]*Node
    head, tail        *Node
    capacity, current int
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        cacheMap: make(map[int]*Node),
        capacity: capacity,
        head:     head,
        tail:     tail,
    }
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cacheMap[key]; exists {
        this.moveToTail(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.cacheMap[key]; exists {
        node.value = value
        this.moveToTail(node)
    } else {
        if this.current == this.capacity {
            this.evict()
        }
        node := &Node{key: key, value: value}
        this.cacheMap[key] = node
        this.addToTail(node)
        this.current++
    }
}

func (this *LRUCache) moveToTail(node *Node) {
    // Remove node from its current position
    node.prev.next = node.next
    node.next.prev = node.prev
    // Add node to tail
    this.addToTail(node)
}

func (this *LRUCache) addToTail(node *Node) {
    node.prev = this.tail.prev
    node.next = this.tail
    this.tail.prev.next = node
    this.tail.prev = node
}

func (this *LRUCache) evict() {
    // Remove the node right after head
    lru := this.head.next
    this.head.next = lru.next
    lru.next.prev = this.head
    delete(this.cacheMap, lru.key)
    this.current--
}
