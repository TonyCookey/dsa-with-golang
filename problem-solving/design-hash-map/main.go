package design_hash_map

type MyHashMap struct {
	hashmap [][]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		hashmap: make([][]int, 1000001),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hashmap[key] = []int{key, value}
}

func (this *MyHashMap) Get(key int) int {
	if this.hashmap[key] != nil {
		return this.hashmap[key][1]
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	if this.Get(key) == -1 {
		return
	} else {
		this.hashmap[key] = nil
	}

}
