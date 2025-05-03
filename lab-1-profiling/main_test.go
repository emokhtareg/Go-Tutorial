package main

import (
	"testing"
)

func BenchmarkSquare(b *testing.B) {

	for i := 0; i < 10; i++ {
		square(i)
	}

}


func BenchmarkSquare2(b *testing.B) {

	for i := 0; i < 10; i++ {
		square2(i)
	}

}