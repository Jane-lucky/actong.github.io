package Alth

import "sort"

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
*/
func ThreeSum(nums []int) [][]int {
	tmp := make([][]int, 0)
	len := len(nums)

	sort.Ints(nums)
	if len < 3 {
		return tmp
	}

	if len == 3 {
		if nums[0]+nums[1]+nums[2] != 0 {
			return tmp
		} else {
			return append(tmp, nums)
		}
	}

	//L,R为移动指针，i为标识符指针
	for i := 0; i < len; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		L := i + 1
		R := len - 1
		for L < R {
			//判断是否与上次提交的相同
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				tmp = append(tmp, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L += 1
				}
				for L < R && nums[R] == nums[R-1] {
					R -= 1
				}
				L += 1
				R -= 1
			} else if sum > 0 {
				R -= 1
			} else {
				L += 1
			}
		}
	}

	return tmp

}

func ThreeSum1(nums []int) [][]int {
	//仅计算左右指针的加是否等于i
	tmp:=[][]int{}

	sort.Ints(nums)

	l:=len(nums)
	if  l<3 {
		return tmp
	}

	for i := 0; i < l; i++ {
		//去重
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		if nums[i]>0 {
			break
		}else{
			start:=-1*nums[i]
			L:=i+1
			R:=l-1
			for L<R{
				if nums[L]+nums[R]==start {
					tmp = append(tmp, []int{nums[i],nums[L],nums[R]})
					//去除双指针中的重复
					for L<R && nums[L]==nums[L+1]{
						L+=1
					}
					for L<R && nums[R]==nums[R-1]{
						R-=1
					}
					L+=1
					R-=1
				} else if nums[L]+nums[R]>start{
					R-=1
				}else{
					L+=1
				}
			}
		}
		
	}
	return tmp
}