# ratelimit [![Build Status](https://travis-ci.org/yangwenmai/ratelimit.png?branch=master)](https://travis-ci.org/yangwenmai/ratelimit) [![Documentation](https://godoc.org/github.com/yangwenmai/ratelimit?status.svg)](http://godoc.org/github.com/yangwenmai/ratelimit) [![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/ratelimit)](https://goreportcard.com/report/github.com/yangwenmai/ratelimit) [![Coverage Status](https://coveralls.io/repos/github/yangwenmai/ratelimit/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/ratelimit?branch=master)

基于令牌桶算法和漏桶算法来实现的限速限流，Golang实现。

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/yangwenmai/ratelimit.svg)](https://starcharts.herokuapp.com/yangwenmai/ratelimit)

# 算法介绍

## 漏桶算法

漏桶算法思路很简单，水（请求）先进入到漏桶里，漏桶以一定的速度出水，当水流入速度过大会直接溢出，可以看出漏桶算法能强行限制数据的传输速率。

![leaky-bucket](https://github.com/yangwenmai/ratelimit/blob/master/doc/leaky-bucket.png)

漏桶算法示意图

漏桶算法可以很好地限制容量池的大小，从而防止流量暴增。

## 令牌桶算法

令牌桶算法的原理是系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务。

![token-bucket](https://github.com/yangwenmai/ratelimit/blob/master/doc/token-bucket.jpg)

令牌桶算法示意图

令牌桶算法通过发放令牌，根据令牌的rate频率做请求频率限制，容量限制等。

### 示例

```go
package main

import (
	"log"
	"time"

	"github.com/yangwenmai/ratelimit/leakybucket"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

func main() {
    // rate limit: simple
	rl := simpleratelimit.New(10, time.Second)

	for i := 0; i < 100; i++ {
		log.Printf("limit result: %v\n", rl.Limit())
	}
	log.Printf("limit result: %v\n", rl.Limit())

    // rate limit: leaky-bucket
	lb := leakybucket.New()
	b, err := lb.Create("leaky_bucket", 10, time.Second)
	if err != nil {
		log.Println(err)
	}
	log.Printf("bucket capacity:%v", b.Capacity())

    // rate limit: token-bucket
}
```

# 参考资料

1. [限流:漏桶算法和令牌桶算法](http://maiyang.github.io/技术/算法/2017/05/28/rate-limit-algorithm)
2. [维基百科:Token_bucket](https://en.wikipedia.org/wiki/Token_bucket)
3. [维基百科:Leaky_bucket](https://en.wikipedia.org/wiki/Leaky_bucket)
4. [接口限流实践](http://www.cnblogs.com/LBSer/p/4083131.html)
5. [流量调整和限流技术](http://colobu.com/2014/11/13/rate-limiting/)
