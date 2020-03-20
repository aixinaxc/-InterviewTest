package algorithm

import (
	"sync"
	"time"
)

type TokenBucket struct {
	tokenInterval   time.Duration //生产token间隔时间
	tokenMaxNum     uint64 //token的最大数量
	tokenSurplusNum uint64 //token的剩余数量
	tokenTicker     *time.Ticker // 定时器 timer
	tokenMutex      *sync.Mutex //token锁
}

//初始化token
func New(tokenInterval time.Duration,tokenMaxNum uint64) (tb *TokenBucket)  {
	tb = & TokenBucket{
		tokenInterval:tokenInterval,
		tokenMaxNum:tokenMaxNum,
		tokenSurplusNum:tokenMaxNum,
		tokenTicker:time.NewTicker(tokenInterval),
		tokenMutex:&sync.Mutex{},
	}
	go tb.productionToken()
	return
}

//创建token
func (tb *TokenBucket) productionToken()  {
	for {
		<- tb.tokenTicker.C
		tb.tokenMutex.Lock()
		if tb.tokenSurplusNum < tb.tokenMaxNum {
			tb.tokenSurplusNum++
		}
		tb.tokenMutex.Unlock()
	}
}

//消费token
func (tb *TokenBucket) consumeToken(useTokenNum uint64) bool {
	tb.tokenMutex.Lock()
	defer tb.tokenMutex.Unlock()
	if useTokenNum <= tb.tokenSurplusNum {
		tb.tokenSurplusNum -= useTokenNum
		return true
	}
	return false
}

//取出token
func (tb *TokenBucket) TakeToken(useTokenNum uint64) bool  {
	return tb.consumeToken(useTokenNum)
}