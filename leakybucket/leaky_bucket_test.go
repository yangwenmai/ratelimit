package leakybucket

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	Convey("start testing", t, func() {
		Convey("create capacity 100", func() {
			s := New()
			bucket, err := s.Create("testCreateBucket", 100, time.Second)
			if err != nil {
				t.Error(err)
			}
			if capacity := bucket.Capacity(); capacity != 100 {
				t.Fatalf("expected capacity of %d, got %d", 100, capacity)
			}
		})
		Convey("same name", func() {
			s := New()
			bucket, err := s.Create("testSame", 10, time.Second)
			if err != nil {
				t.Error(err)
			}
			bucket2, err := s.Create("testSame", 100, time.Second)
			if err != nil {
				t.Error(err)
			}
			if bucket != bucket2 {
				t.Fatalf("not same.")
			}
		})
		Convey("add()", func() {
			s := New()
			bucket, err := s.Create("testAdd", 100, time.Second)
			if err != nil {
				t.Fatal(err)
			}

			addAndTestRemaining := func(add, remaining uint) {
				if state, err := bucket.Add(add); err != nil {
					t.Fatal(err)
				} else if bucket.Remaining() != state.Remaining {
					t.Fatalf("expected bucket and state remaining to match, bucket is %d, state is %d",
						bucket.Remaining(), state.Remaining)
				} else if state.Remaining != remaining {
					t.Fatalf("expected %d remaining, got %d", remaining, state.Remaining)
				}
			}
			addAndTestRemaining(1, 99)
			addAndTestRemaining(3, 96)
			addAndTestRemaining(90, 6)
			if _, err := bucket.Add(7); err == nil {
				t.Fatalf("expected ErrorFull, received no error")
			} else if err != ErrorFull {
				t.Fatalf("expected ErrorFull, received %v", err)
			}
		})
		Convey("Reset()", func() {
			s := New()
			bucket, err := s.Create("testReset", 1, time.Millisecond)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := bucket.Add(1); err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Millisecond * 2)
			bucket.Reset()
			if bucket.Remaining() != 1 {
				t.Fatalf("expected 1, got %d", bucket.Remaining())
			}
			state, err := bucket.Add(1)
			if err != nil {
				t.Fatal(err)
			}
			if state.Remaining != 0 {
				t.Fatalf("expected full bucket, got %d", state.Remaining)
			}
			if state.Reset.Unix() < time.Now().Unix() {
				t.Fatalf("reset time is in the past")
			}
			if bucket.Reset().Unix() < time.Now().Unix() {
				t.Fatalf("reset time is in the past")
			}
		})
	})
}

func BenchmarkLimit(b *testing.B) {
	Convey("start testing benchmark", b, func() {
	})
}
