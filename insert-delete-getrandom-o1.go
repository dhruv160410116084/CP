package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set_map map[int]int
	set     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set_map[val]; ok {
		return false
	}

	this.set = append(this.set, val)

	this.set_map[val] = len(this.set) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// fmt.Println(val, this.set_map[val])
	if i, ok := this.set_map[val]; ok {
		fmt.Println("in if")
		fmt.Println(this.set, this.set_map)

		delete(this.set_map, val)
		if len(this.set_map) > 0 {
			this.set[i] = this.set[len(this.set)-1]
			this.set_map[this.set[i]] = i
			this.set = this.set[:len(this.set)-1]
		} else {
			this.set = this.set[:0]
		}
		fmt.Println(this.set, this.set_map)
		return true

	}
	return false
}

func (this *RandomizedSet) GetRandom() int {

	r := rand.Intn(len(this.set))
	return this.set[r]
	// fmt.Println(r)
	// return this.set_map[r]
}

func main_insert_delete_getrandom() {
	obj := Constructor()
	// fmt.Println(obj.set_map)
	// fmt.Println(obj.Insert(0))
	// fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(-1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)
	fmt.Println(obj.GetRandom(), obj.set_map, obj.set)

	// fmt.Println(obj.Remove(0))

	// fmt.Println(obj.GetRandom(), obj.set_map, obj.set)

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
