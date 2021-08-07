package main

import (
	"fmt"
	"strconv"
	"strings"
)

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func CountNumberAppears(arr []int) {
	if len(arr) < 1 {
		fmt.Println(1)
		return
	}
	var count int
	var total []string
	for j := 0; j < len(arr); j++ {
		count = 0
		for i := 0; i < len(arr); i++ {
			if arr[j] == arr[i] {
				count += 1
				if len(total) > j+1 {
					total[j] = strconv.Itoa(arr[j]) + " : " + strconv.Itoa(count)
				} else {
					total = append(total, strconv.Itoa(arr[j])+" : "+strconv.Itoa(count))
				}

			}
		}
	}

	stringNumbers := "number --> total\n"
	total = unique(total)
	for _, val2 := range total {
		stringSplit := strings.Split(val2, " : ")
		stringNumbers += stringSplit[0] + "-->" + stringSplit[1] + "\n"

	}
	fmt.Println(stringNumbers)
}

func main() {
	CountNumberAppears([]int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6})
}
