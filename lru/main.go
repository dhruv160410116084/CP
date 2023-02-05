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

func main() {
	fmt.Println("hello")
}
