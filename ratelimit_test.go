package ratelimit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	Convey("start test Limit(), result=false,", t, func() {
		rl := New()
		So(rl.Limit(), ShouldEqual, false)
	})
}

func BenchmarkLimit(b *testing.B) {
	Convey("start BenchmarkLimit", b, func() {
		rl := New()
		for i := 0; i < b.N; i++ {
			rl.Limit()
		}
	})
}
