package main

import (
	"testing"
)

func BenchmarkAddToSliceUnique(b *testing.B) {
	strings := []string{
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
	}
	for i := 0; i < b.N; i++ {
		for _, v := range strings {
			AddToSliceUnique(v)
		}
	}
}

func BenchmarkAddToMapUnique(b *testing.B) {
	strings := []string{
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
		"abcdefgh", "ijklmnop", "qrstuvwx", "yzabcdef", "ghijklmn",
		"opqrstuv", "wxyzabcd", "efghijkl", "mnopqrst", "uvwxyzab",
	}
	for i := 0; i < b.N; i++ {
		for _, v := range strings {
			AddToMapUnique(v)
		}
	}
}
