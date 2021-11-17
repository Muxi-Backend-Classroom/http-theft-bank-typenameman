package main

import (
	"fmt"
)

func change(arr []int, i int, j int) {
	var mid int
	mid = arr[i]
	arr[i] = arr[j]
	arr[j] = mid
}

var t = 0

func re(n int) int {
	num := 1
	for i := n; i >= 1; i-- {
		num = num * i
	}
	return num
}
func ReturnFunc(array []int, num [][]int, p int) {
	if p == len(array)-1 {
		copy(num[t], array) //此处如果使用append则会自动销毁内部切片(我也不明白为什么)，只得使用此替代方案
		t++
	} else {
		for i := p; i <= len(array)-1; i++ {
			var arr = make([]int, len(array))
			copy(arr, array)
			change(arr, i, p)
			ReturnFunc(arr, num, p+1)
		}
	}
}
func permute(nums []int) [][]int {
	var res [][]int
	for i := 0; i <= re(n)-1; i++ {
		var a = make([]int, n)
		res = append(res, a)
	}
	ReturnFunc(nums, res, 0)
	return res
}

var n int

func main() {
	fmt.Scanf("%d", &n)
	testSlice := make([]int, n)
	for i := 0; i <= n-1; i++ {
		fmt.Scan(&testSlice[i])
	}
	res := permute(testSlice)
	fmt.Println(res)
}
