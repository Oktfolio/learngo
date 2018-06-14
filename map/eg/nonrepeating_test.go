package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"bbb", 1},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"测试中文", 4},
		{"一二三一二三一二三", 3},
		{"黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞", 8},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s);
			actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞", 8

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s);
			actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
