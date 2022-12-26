package helpers

import (
	"fmt"
	"math"
	"strconv"
)

func StrToI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StrSliceToIntSlice(strSlice []string) []int {
	var iSlice []int

	for _, s := range strSlice {
		iSlice = append(iSlice, StrToI(s))
	}

	return iSlice
}

func IntSliceToStrSlice(intSlice []int) []string {
	var iSlice []string

	for _, s := range intSlice {
		iSlice = append(iSlice, fmt.Sprint(s))
	}

	return iSlice
}

func IMin(x int, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func IMax(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func Min[V int | float64](x V, y V) V {
	return V(math.Min(float64(x), float64(y)))
}

func Max[V int | float64](x V, y V) V {
	return V(math.Max(float64(x), float64(y)))
}
