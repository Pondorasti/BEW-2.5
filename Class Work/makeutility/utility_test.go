package main

import (
	"testing"
)

// func TestLookUp(t *testing.T) {
// files := LookUp("./")
// assert.Equal(t, len(files), 7)
// }

func BenchmarkLookUp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LookUp("./")
	}
}
