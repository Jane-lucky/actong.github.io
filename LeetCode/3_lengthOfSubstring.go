package Alth

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

思路：
1.如果字符串的长度为0，返回0
2.循环遍历字符串
map[byte]int 存储字符串出现的次数
slen 记录最长子串的长度

中间变量count

*/
func LengthOfLongestSubstring(s string) int {
	var slen int
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		count := 0
		smap := make(map[byte]int)
		for j := i; j < len(s); j++ {
			if _, ok := smap[s[j]]; !ok {
				smap[s[j]]++
				count++
			} else {
				break
			}
		}

		if count > slen {
			slen = count
		}
	}
	return slen
}
