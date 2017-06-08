package ratelimit

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	Convey("初始化一个限制，然后进行测试限速: ", t, func() {
		rl := New(1, time.Second)
		for i := 0; i < 10; i++ {
			if i == 0 {
				So(rl.Limit(), ShouldEqual, false)
			} else {
				So(rl.Limit(), ShouldEqual, true)
			}
		}
	})
	Convey("第一次进行限速后，每次限速之前再重置，检查限速是否通过: ", t, func() {
		rl := New(1, time.Second)
		for i := 0; i < 10; i++ {
			if i == 0 {
				So(rl.Limit(), ShouldEqual, false)
			} else {
				rl.Undo()
				So(rl.Limit(), ShouldEqual, false)
			}
		}
	})
}

func BenchmarkLimit(b *testing.B) {
	Convey("开始压测限速", b, func() {
		rl := New(1, time.Second)
		for i := 0; i < b.N; i++ {
			rl.Limit()
		}
	})
}
