package Alth

import (
	"strings"
)

/**
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。


思路：
1. 判断字符串s的长度和words中字符的长度，小于，return

2.截取与words等长度的字符串s

*/

func FindSubstring(s string, words []string) []int {
	//words中存在相同的元素时没办法判断
	//slen字符串长度 wlen words单字符的长度
	slen := len(s)
	wlen := len(words[0])

	k := len(words)

	res := []int{}
	if slen < wlen*k {
		return res
	}

	for i := 0; i < slen-wlen*k+1; i++ {
		tag := 0
		for _, v := range words {
			start := i
			tmp := s[start : start+wlen*k]
			if !strings.Contains(tmp, v) {
				tag = 1
				break
			}
		}
		if tag == 0 {
			res = append(res, i)
		}

	}
	return res

}

//双map解决
func FindSubstring1(s string, words []string) []int {
	//采用map的方式存储words

	//slen字符串长度 wlen words单字符的长度
	slen := len(s)
	wlen := len(words[0])

	k := len(words)

	res := []int{}
	if slen < wlen*k {
		return res
	}

	mmap := make(map[string]int)

	for i := 0; i < len(words); i++ {
		mmap[words[i]]++
	}

	//遍历子串
	for i := 0; i < slen-wlen*k+1; i++ {

		mmap1 := make(map[string]int)
		var count int
		for j := 0; j < k; j++ {
			start := i + wlen*j
			t := s[start : start+wlen]
			if num, ok := mmap[t]; ok && num > mmap1[t] {
				mmap1[t]++
				count++
			} else {
				break
			}
		}
		if count == k {
			res = append(res, i)
		}

	}
	return res
}

//滑动窗口
func FindSubstring2(s string, words []string) []int {
	return nil
}
