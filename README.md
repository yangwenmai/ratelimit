# ratelimit [![Build Status](https://travis-ci.org/yangwenmai/ratelimit.png?branch=master)](https://travis-ci.org/yangwenmai/ratelimit) [![Documentation](https://godoc.org/github.com/yangwenmai/ratelimit?status.svg)](http://godoc.org/github.com/yangwenmai/ratelimit)

基于令牌桶算法和漏桶算法来实现的限速限流，Golang实现。

# 算法介绍

## 漏桶算法

漏桶算法思路很简单，水（请求）先进入到漏桶里，漏桶以一定的速度出水，当水流入速度过大会直接溢出，可以看出漏桶算法能强行限制数据的传输速率。
![leaky-bucket](https://raw.githubusercontent.com/yangwenmai/ratelimit/master/leaky_bucket.png)
漏桶算法可以很好地限制容量池的大小，从而防止流量暴增。

## 令牌桶算法

令牌桶算法的原理是系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务。
![token-bucket](https://raw.githubusercontent.com/yangwenmai/ratelimit/master/token_bucket.jpg)
令牌桶算法通过发放令牌，根据令牌的rate频率做请求频率限制，容量限制等。

# 参考资料

1. [限流:漏桶算法和令牌桶算法](http://maiyang.github.io/技术/算法/2017/05/28/rate-limit-algorithm)
