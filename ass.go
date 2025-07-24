package ass

import (
	"fmt"
	"strings"
)

func A0(e error) {
	if e != nil {
		panic(e)
	}
}

func A1[T1 any](v1 T1, e error) T1 {
	if e != nil {
		panic(e)
	}
	return v1
}

func A2[T1, T2 any](v1 T1, v2 T2, e error) (T1, T2) {
	if e != nil {
		panic(e)
	}
	return v1, v2
}

func A3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, e error) (T1, T2, T3) {
	if e != nil {
		panic(e)
	}
	return v1, v2, v3
}

// Functions should never return more than 3 values, are you insane? Go make a struct and find God...

func repeat(n, m int) []any {
	s := make([]any, m)
	for i := range s {
		s[i] = n
	}
	return s
}

func writeRepeat(sb *strings.Builder, n int, fmtString, separator, suffix string) {
	numPlaceholders := strings.Count(fmtString, "%d")
	for i := range n {
		fmt.Fprintf(sb, fmtString, repeat(i+1, numPlaceholders)...)
		if i < n-1 {
			sb.WriteString(separator)
		}
	}
	sb.WriteString(suffix)
}

func AN(n int) (string, error) {
	if n < 0 {
		return "ASSHOLE", fmt.Errorf("n must be >= 0")
	} else if n == 0 {
		return "func A0(e error){if e!=nil{panic(e)}}", nil
	} else {
		sb := &strings.Builder{}
		fmt.Fprintf(sb, "func A%d[", n)
		writeRepeat(sb, n, "T%d", ",", " any](")
		writeRepeat(sb, n, "v%d T%d", ",", ",e error)(")
		writeRepeat(sb, n, "T%d", ",", "){if e!=nil{panic(e)};return ")
		writeRepeat(sb, n, "v%d", ",", "}")
		return sb.String(), nil
	}
}
