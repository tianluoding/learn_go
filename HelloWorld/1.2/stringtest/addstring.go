package stringtest

import "strings"

func AddString() string {
	var s, sep string
	for i := 0; i < 10; i++ {
		s += sep + "test"
		sep = " "
	}
	return s
}

func JoinString() string {
	var s string
	strs := [10]string{"test", "test", "test", "test", "test", "test", "test", "test", "test", "test"}
	s = strings.Join(strs[:], " ")
	return s
}
