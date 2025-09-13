package utils

import "strconv"

func ParseUint(s string) uint {
    i, _ := strconv.ParseUint(s, 10, 64)
    return uint(i)
}
