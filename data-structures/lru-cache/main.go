package lru_cache

type LRUCache struct {
	Cache    map[int]*Node
	Capacity int
	Head     *Node
	Tail     *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	Head := &Node{
		Key: 0,
		Val: 0,
	}
	Tail := &Node{
		Key: -1,
		Val: -1,
	}
	Tail.Prev, Head.Next = Head, Tail

	return LRUCache{
		Cache:    make(map[int]*Node),
		Capacity: capacity,
		Head:     Head,
		Tail:     Tail,
	}
}
func (this *LRUCache) insert(node *Node) {
	this.Cache[node.Key] = node
	MRU := this.Tail.Prev
	MRU.Next, node.Prev = node, MRU
	this.Tail.Prev, node.Next = node, this.Tail
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.Prev, node.Next

	prev.Next, next.Prev = next, prev
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.Cache[key]; exists {
		this.remove(node)
		this.insert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.Cache[key]; exists {
		this.remove(node)
	}
	this.insert(&Node{
		Key: key,
		Val: value,
	})

	if len(this.Cache) > this.Capacity {
		lru := this.Head.Next
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}
}
