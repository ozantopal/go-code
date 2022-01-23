package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}

func isIsomorphic(s string, t string) bool {
	sList := []byte(s)
	tList := []byte(t)

	if (s == "" && t != "") || (len(sList) != len(tList)) {
		return false
	}

	sMap := map[byte]byte{}
	pMap := map[byte]byte{}
	for index, b := range sList {
		if _, ok := sMap[b]; !ok {
			if _, ok = pMap[tList[index]]; !ok {
				sMap[b] = tList[index]
				pMap[tList[index]] = b
			} else {
				if pMap[tList[index]] != b {
					return false
				}
			}
		} else {
			if sMap[b] != tList[index] {
				return false
			}
		}
	}
	return true
}
