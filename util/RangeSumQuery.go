package util

type NumArray struct {
    Elements []int 
}


func Constructor(nums []int) NumArray {
    numArr := NumArray{
			Elements: nums,
		}
		return numArr
}


func (this *NumArray) SumRange(left int, right int) int {
		var res int
		for _, v := range this.Elements[left:right+1] {
			res += v
		}
    return res
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */