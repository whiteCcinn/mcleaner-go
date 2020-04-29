package app

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v7"
	mconfig "mcleaner/config"
	"strconv"
)

// 先驱者主要是处理从kafka拉取数据到redis到过程
// 核心方法主要有一个handle()
var pioneer = new(emptyPioneer)

// 先驱者接口方法
type PioneerInterface interface {
	Handle(ctx context.Context)
}

// 先驱者实现结构体
type emptyPioneer struct{}

// 用于创建返回全局唯一到先驱者对象
func PioneerBackground() PioneerInterface {
	return pioneer
}

// 核心业务方法
func (p *emptyPioneer) Handle(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_9_0_0 // specify appropriate version
	config.Consumer.Return.Errors = true

	// 捕捉panic错误
	errorChan := make(chan interface{})
	defer func() {
		if r := recover(); r != nil {
			errorChan <- r
		}
	}()

	group, err := sarama.NewConsumerGroup(mconfig.Config.Kafka.Consumer.Broker, mconfig.Config.Kafka.Consumer.Group, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	//go func() {
	// Iterate over consumer sessions.
	consumerCtx := context.Background()
	for {
		topics := []string{"my-topic"}
		handler := consumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(consumerCtx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
	//}()
}

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// 处理业务逻辑
	client := redis.NewClient(&redis.Options{
		Addr:     mconfig.Config.Redis.Store.Ip.(string) + ":" + strconv.Itoa(int(mconfig.Config.Redis.Store.Port)),
		Password: mconfig.Config.Redis.Store.Auth,
		DB:       int(mconfig.Config.Redis.Store.Db),
	})

	// 如果有消息，会从"消息链"的channel中读取
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Printf("%s", msg.Value)

		// partition + md5(消息内容)
		key := strconv.Itoa(int(msg.Partition)) + fmt.Sprintf("%x", md5.Sum(msg.Value))

		// 1. 存储数据
		err := client.Set(key, fmt.Sprintf("%s", msg.Value), 0).Err()
		if err != nil {
			panic(err)
		}

		// 2. 存储数据的key到list
		listKey := fmt.Sprintf("%s:%d:list-key", mconfig.Config.App.Name, mconfig.Config.App.Id)
		err = client.LPush(listKey, key).Err()
		if err != nil {
			panic(err)
		}

		// todo：标记offset，等客户端自动commit offset
		// 这里我们后面需要改成我们自己手动管理offset
		sess.MarkMessage(msg, "")
	}

	return nil
}
