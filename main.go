package main

import (
	"log"
	"time"

	"github.com/yangwenmai/ratelimit/leakybucket"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

func main() {
	rl := simpleratelimit.New(10, time.Second)

	for i := 0; i < 100; i++ {
		rl.Limit()
	}
	log.Printf("limit: %v", rl.Limit())
	lb := leakybucket.New()
	b, err := lb.Create("leaky_bucket", 10, time.Second)
	if err != nil {
		log.Println(err)
	}
	log.Printf("capacity:%v", b.Capacity())
}
