// Package insertdeletegetrandomo1 solves LeetCode 380. Insert Delete Getrandom O1.
// https://leetcode.com/problems/insert-delete-getrandom-o1/
package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	valToIndex map[int]int
	nums       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int, 8),
		nums:       make([]int, 0, 8),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, existed := this.valToIndex[val]
	if !existed {
		this.nums = append(this.nums, val)
		this.valToIndex[val] = len(this.nums) - 1
	}
	return !existed
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, existed := this.valToIndex[val]
	if existed {
		lastestVal := this.nums[len(this.nums)-1]
		this.nums[idx] = lastestVal
		this.valToIndex[lastestVal] = idx

		delete(this.valToIndex, val)
		this.nums = this.nums[:len(this.nums)-1]
	}
	return existed
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Int() % len(this.nums)
	return this.nums[idx]
}
