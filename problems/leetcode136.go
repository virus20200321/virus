package problems

//异或也叫半加运算，其运算法则相当于不带进位的二进制加法：二进制下用1表示真，0表示假，
//则异或的运算法则为：0⊕0=0，1⊕0=1，0⊕1=1，1⊕1=0（同为0，异为1），
//这些法则与加法是相同的，只是不带进位，所以异或常被认作不进位加法。
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] ^ nums[i-1]
	}
	return nums[len(nums)-1]
}

func singleNumber2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] ^ nums[i]

	}
	return nums[len(nums)-1]
}
