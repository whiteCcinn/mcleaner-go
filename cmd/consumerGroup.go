/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd
//
//import (
//	"context"
//	"fmt"
//	"github.com/Shopify/sarama"
//	"github.com/spf13/cobra"
//)
//
////type exampleConsumerGroupHandler struct{}
////
////func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
////func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
////func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
////	for msg := range claim.Messages() {
////		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
////		fmt.Printf("%s", msg.Value)
////
////		//// 处理业务逻辑
////		client := redis.NewClient(&redis.Options{
////			Addr:     "localhost:6379",
////			Password: "", // no password set
////			DB:       0,  // use default DB
////		})
////
////		// partition + md5(消息内容)
////		key := strconv.Itoa(int(msg.Partition)) + fmt.Sprintf("%x", md5.Sum(msg.Value))
////
////		// 1. 存储数据
////		err := client.Set(key, string(msg.Value[:]), 0).Err()
////		if err != nil {
////			panic(err)
////		}
////
////		// 2. 存储数据的key到list
////		listKey := fmt.Sprintf("%s:%d:list-key", config.Config.App.Name, config.Config.App.Id)
////		err = client.LPush(listKey, key).Err()
////		if err != nil {
////			panic(err)
////		}
////
////		sess.MarkMessage(msg, "")
////	}
////	return nil
////}
//
//// consumerGroupCmd represents the consumerGroup command
//var consumerGroupCmd = &cobra.Command{
//	Use:   "consumerGroup",
//	Short: "消费者组",
//	Long: `A longer description that spans multiple lines and likely contains examples
//and usage of using your command. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		config := sarama.NewConfig()
//		config.Version = sarama.V0_9_0_0 // specify appropriate version
//		config.Consumer.Return.Errors = true
//
//		group, err := sarama.NewConsumerGroup([]string{"mkafka1:9092"}, "my-group", config)
//		if err != nil {
//			panic(err)
//		}
//		defer func() { _ = group.Close() }()
//
//		// Track errors
//		go func() {
//			for err := range group.Errors() {
//				fmt.Println("ERROR", err)
//			}
//		}()
//
//		// Iterate over consumer sessions.
//		ctx := context.Background()
//		for {
//			topics := []string{"my-topic"}
//			handler := exampleConsumerGroupHandler{}
//
//			// `Consume` should be called inside an infinite loop, when a
//			// server-side rebalance happens, the consumer session will need to be
//			// recreated to get the new claims
//			err := group.Consume(ctx, topics, handler)
//			if err != nil {
//				panic(err)
//			}
//		}
//	},
//}

//func init() {
//	rootCmd.AddCommand(consumerGroupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerGroupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerGroupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
