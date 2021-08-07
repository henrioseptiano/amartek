package main

import "fmt"

func stringMerge(string1, string2 string) string {
	var resString string
	if len(string2) < len(string1) {
		for i := 0; i < len(string1); i++ {
			if len(string2) <= i {
				resString += string(string1[i])
			} else {
				for j := i; j < len(string2); j++ {
					resString += string(string1[i]) + string(string2[j])
					break
				}
			}
		}
	} else {
		for i := 0; i < len(string2); i++ {
			if len(string1) <= i {
				resString += string(string2[i])
			} else {
				for j := i; j < len(string1); j++ {
					resString += string(string1[i]) + string(string2[j])
					break
				}
			}
		}
	}
	return resString
}

func main() {
	fmt.Println(stringMerge("saya", "kamu"))
	fmt.Println(stringMerge("kosong", "ada"))
	fmt.Println(stringMerge("ada", "kosong"))
}
