package simpleratelimit

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	Convey("start testing New(0, 0)", t, func() {
		rl := New(0, 0*time.Second)
		So(rl.Limit(), ShouldEqual, false)
	})
	Convey("start testing limit", t, func() {
		rl := New(1, time.Second)
		for i := 0; i < 10; i++ {
			if i == 0 {
				So(rl.Limit(), ShouldEqual, false)
			} else {
				So(rl.Limit(), ShouldEqual, true)
			}
		}
	})
	Convey("start testing updateRate", t, func() {
		rl := New(1, time.Second)
		for i := 0; i < 2; i++ {
			if i == 0 {
				So(rl.Limit(), ShouldEqual, false)
			} else {
				rl.UpdateRate(2)
				So(rl.Limit(), ShouldEqual, true)
			}
		}
	})
	Convey("start testing undo", t, func() {
		rl := New(1, time.Second)
		for i := 0; i < 10; i++ {
			if i == 0 {
				So(rl.Limit(), ShouldEqual, false)
			} else {
				rl.Undo()
				So(rl.Limit(), ShouldEqual, false)
			}
		}
		for i := 0; i < 10; i++ {
			rl.Limit()
		}
		rl.Undo()
	})
}

func BenchmarkLimit(b *testing.B) {
	Convey("start testing benchmark", b, func() {
		rl := New(1, time.Second)
		for i := 0; i < b.N; i++ {
			rl.Limit()
		}
	})
}
