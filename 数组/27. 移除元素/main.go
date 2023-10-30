package main

func main() {

}

func removeElement(nums []int, val int) int {
	res := 0
	for i := range nums {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
			continue
		}
	}
	return res
}
