package piscine

import "strings"

func StrLen(str string) int {
	a := strings.Count(str, "") -1
    return a
}
