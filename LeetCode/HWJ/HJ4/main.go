package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	arrstr := SplitStr(str)
	for _, v := range arrstr {
		fmt.Println(v)
	}
}

func SplitStr(str string) []string {
	llen := len(str)
	res := []string{}
	if llen < 8 {
		s := ""
		for i := 0; i < 8-llen; i++ {
			s += "0"
		}

		res = append(res, str+s)
		return res
	}
	for i := 0; i < llen; i += 8 {
		stmp := ""
		if i+8 >= llen {
			stmp = str[i:]
		} else {
			stmp = str[i : i+8]
		}

		if len(stmp) == 8 {
			res = append(res, stmp)
		} else {
			s := ""
			for i := 0; i < 8-len(stmp); i++ {
				s += "0"
			}

			res = append(res, stmp+s)
		}

	}
	return res
}
