package app

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v7"
	mconfig "mcleaner/config"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 先驱者主要是处理从kafka拉取数据到redis到过程
// 核心方法主要有一个handle()
var sinker = new(emptySinker)

// 先驱者接口方法
type SinkerInterface interface {
	Handle(ctx context.Context)
}

// 先驱者实现结构体
type emptySinker struct{}

// 用于创建返回全局唯一到先驱者对象
func SinkerBackground() SinkerInterface {
	return sinker
}

// 核心业务方法
func (p *emptySinker) Handle(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_9_0_0 // specify appropriate version
	config.Consumer.Return.Errors = true

	// 处理业务逻辑
	client := redis.NewClient(&redis.Options{
		Addr:     mconfig.Config.Redis.Store.Ip.(string) + ":" + strconv.Itoa(int(mconfig.Config.Redis.Store.Port)),
		Password: mconfig.Config.Redis.Store.Auth,
		DB:       int(mconfig.Config.Redis.Store.Db),
	})

	listKey := fmt.Sprintf("%s:%d:list-key", mconfig.Config.App.Name, mconfig.Config.App.Id)

	interval := 5
	ticker := time.NewTicker(time.Second * 5)
	defer func() { ticker.Stop() }()

	rw := sync.RWMutex{}
	wg := sync.WaitGroup{}

	i := 0
	wg.Add(1)
	go func() {
		defer func() { wg.Done() }()
		for {
			select {
			case <-ticker.C:
				// 读锁
				rw.RLock()
				fmt.Printf("core: %d, tps: %d\n", runtime.NumCPU(), i/(interval))
				rw.RUnlock()

				// 写锁
				rw.Lock()
				i = 0
				rw.Unlock()
			}
		}
	}()

	for count := 0; count < 9; {
		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			for {
				result, err := client.BLPop(30*time.Second, listKey).Result()
				if err != nil {
					println("sinker 超时")
					continue
				}

				dataKey := result[1]

				// 获取数据
				//msg, err := client.Get(dataKey).Result()
				_, err = client.Get(dataKey).Result()

				// 消费数据
				//fmt.Printf("消费了数据: %s\n", fmt.Sprintf("%s", msg))

				// 从redis移除，完成redis的ack
				client.Del(dataKey)

				// 写锁
				rw.Lock()
				i++
				rw.Unlock()
			}
		}()
		count++
	}

	wg.Wait()
}
