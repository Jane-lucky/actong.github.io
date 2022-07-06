package Alth

import (
	"sort"
)

/**
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次
样例：
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

思路:
字符串排序

将已经提供的字符串数组进行排序，因为所包含的字符串相等

map[string][]string{}
判断排序
*/
func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	mmap := map[string][]string{}

	for _, v := range strs {
		byteArry := []byte(v)

		sort.Slice(byteArry, func(k, j int) bool {
			return byteArry[k] < byteArry[j]
		})
		mmap[string(byteArry)] = append(mmap[string(byteArry)], v)
	}

	t := [][]string{}

	for _, v := range mmap {
		t = append(t, v)
	}

	return t

}

func GroupAnagrams1(strs []string) [][]string {
	mp := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

		sotesStr := string(s)
		mp[sotesStr] = append(mp[sotesStr], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
