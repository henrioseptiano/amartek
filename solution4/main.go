package main

import "fmt"

func lexicograhpicallyMinimalString(lines int, string1 string, string2 string) string {
	var finalString string
	var count int
	for i := 0; i < len(string1); i++ {
		if count < lines {
			for j := 0; j < len(string2); j++ {
				if string1[i] == string2[i] {
					finalString += string(string1[i]) + string(string2[j])
					finalString += string1[i+1:i+lines+1] + string2[j+1:j+lines+1]
					count = i + lines + 2
					break
				} else if string1[i] > string2[j] {
					finalString += string(string2[j])
				} else if string1[i] < string2[j] {
					finalString += string(string1[i])
					break
				}
				count++
			}
		}
	}
	//fmt.Println(finalString)
	if count > lines {
		finalString += string1[count-1:] + string2[count-1:]
	} else {
		finalString += string1[count-1:] + string2[count:]
	}
	return finalString
}

func main() {
	fmt.Println(lexicograhpicallyMinimalString(2, "JACK", "DANIEL"))
	fmt.Println(lexicograhpicallyMinimalString(2, "ABACABA", "ABACABA"))
}
