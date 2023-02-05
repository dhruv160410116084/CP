package main

import (
	"container/list"
	"fmt"
)

type Element struct {
	key   int
	value int
}

type MNode struct {
	add   *list.Element
	value int
}

type LRUCache struct {
	addMap map[int]MNode
	dll    *list.List
	cap    int
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{make(map[int]MNode), list.New(), capacity}
}

func (this *LRUCache) Get(key int) int {
	//if not exist return -1

	if _, ok := this.addMap[key]; !ok {
		return -1
	} else {
		this.dll.Remove(this.addMap[key].add)
		// fmt.Println(n)
		e := this.dll.PushFront(Element{key, this.addMap[key].value})
		// fmt.Println(e)
		this.addMap[key] = MNode{e, this.addMap[key].value}
		// fmt.Println(this.addMap[key])

		return this.addMap[key].value

	}

}

func (this *LRUCache) Put(key int, value int) {

	//if capacity is full delete least recently used
	// fmt.Println("len:", this.dll.Len())
	// fmt.Println("list: ", &this.dll)
	_, isExist := this.addMap[key]

	if len(this.addMap) >= this.cap && !isExist {
		// fmt.Println(this.dll)

		// fmt.Println("front: ", this.dll.Front())
		// fmt.Println("back", this.dll.Back().Prev())

		// fmt.Println(this.dll.Back().Prev().Value)
		var e = this.dll.Remove(this.dll.Back())
		fmt.Println("delete:", e)
		delete(this.addMap, e.(Element).key)

	}

	// if not in map insert after head

	if _, ok := this.addMap[key]; !ok {
		e := this.dll.PushFront(Element{key, value})
		// fmt.Println("address of ", e)
		// fmt.Println("head: ", this.dll, "back: ", this.dll.Back(), "front: ", this.dll.Front())
		this.addMap[key] = MNode{e, value}
	} else {
		// else move node to next after head and update address in map
		fmt.Println("address", this.addMap[key].add)
		this.dll.Remove(this.addMap[key].add)
		// fmt.Println(e)
		// fmt.Println("address of ", e)
		// fmt.Println("head: ", this.dll, "back: ", this.dll.Back(), "front: ", this.dll.Front())

		n := this.dll.PushFront(Element{key, value})
		this.addMap[key] = MNode{n, value}
	}

}

func main_lru() {
	q := []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"}
	o := [][]int{{2}, {2}, {2, 6}, {1}, {1, 5}, {1, 2}, {1}, {2}}
	// ans := make([]int,)
	var cache LRUCache

	for i, v := range q {
		if v == "LRUCache" {
			cache = ConstructorLRU(o[i][0])
		} else if v == "put" {
			cache.Put(o[i][0], o[i][1])
			// cache.iForward()
			// cache.iBackward()
		} else if v == "get" {
			fmt.Println("ans:", cache.Get(o[i][0]))
			// cache.iForward()
			// cache.iBackward()
		}
	}

	// fmt.Println(cache)

}
