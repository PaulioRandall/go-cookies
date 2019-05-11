package pkg

import (
	"strconv"
)

// AppendIfEmpty appends m to r if s is empty.
func AppendIfEmpty(s string, r []string, m string) []string {
	if s == "" {
		return append(r, m)
	}
	return r
}

// AppendIfNotUint appends m to r if s is NOT a positive integer.
func AppendIfNotUint(s string, r []string, m string) []string {
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		return append(r, m)
	}
	return r
}

// AppendIfNotUintCSV appends m to r if s is not a CSV of positive integers.
func AppendIfNotUintCSV(s string, r []string, m string) []string {
	if IsUintCSV(s) {
		return r
	}
	return append(r, m)
}
