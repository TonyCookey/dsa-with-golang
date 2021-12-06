package main

import (
	"fmt"
)

type hashTable struct {
	data map[int][]string
}

func (h hashTable) hash(key string) int {
	var hash int
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return hash
}
func (h hashTable) set(key, value string) {
	hash := h.hash(key)
	if _, ok := h.data[hash]; ok {
		h.data[hash] = append(h.data[hash], key, value)
	} else {
		bucket := []string{key, value}
		h.data[hash] = bucket
	}

}
func (h hashTable) get(key string) string {
	hash := h.hash(key)
	if _, ok := h.data[hash]; ok {
		index := len(h.data[hash]) - 1
		return h.data[hash][index]
	} else {
		return h.data[hash][1]
	}

}
func main() {
	data := make(map[int][]string)
	hashtable := hashTable{
		data: data,
	}
	hashtable.set("Tony", "StarBoy")
	hashtable.set("Tony", "Cookey")
	hashtable.set("Tony", "Ogbe")
	hashtable.set("C", "ORORO")
	fmt.Println(hashtable.get("Tony"))

}
