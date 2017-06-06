package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	Convey("start test limit, result=true,", t, func() {
		fmt.Println("ok.")
	})
}

func BenchmarkLimit(b *testing.B) {
	Convey("start BenchmarkLimit", b, func() {
		for i := 0; i < b.N; i++ {
			fmt.Println("BenchmarkLimit.")
		}
	})
}
