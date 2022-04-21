package design_hashset

type MyHashSet struct {
	store []bool
}

func Constructor() MyHashSet {
	newMap := make([]bool, 1000001)
	return MyHashSet{
		store: newMap,
	}
}

func (this *MyHashSet) Add(key int) {
	this.store[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.store[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	if this.store[key] {
		return true
	}
	return false
}
