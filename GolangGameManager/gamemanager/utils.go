package gamemanager

import (
	"strconv"
	"strings"
)

func deleteElements(array []int, elements []int) []int {
	for _, element := range elements {
		for i, val := range array {
			if val == element {
				array = append(array[:i], array[i+1:]...)
				break
			}
		}
	}
	return array
}

func arrayToString(arr []int) string {

	strSlice := make([]string, len(arr))

	for i, num := range arr {
		strSlice[i] = strconv.Itoa(num)
	}

	result := strings.Join(strSlice, " ")
	return result
}
