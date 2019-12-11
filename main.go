package main

import "fmt"

func bubbleDown (index int, nums []int) {
	for i := index; i < len(nums)-1; i++ {
		if nums[i+1] == 0 {
			break
		}
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

func moveZeroes(nums []int)  {

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			bubbleDown(i, nums)
		}
	}

}

func main() {
	arr := []int {0,1,0,3,12}
	moveZeroes(arr)

	fmt.Printf("moved zeros %v\n", arr)
}
