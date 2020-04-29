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

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

// kafkaProducerCmd represents the kafkaProducer command
var kafkaProducerCmd = &cobra.Command{
	Use:   "kafkaProducer",
	Short: "生产消息",
	Long:  `生产消息`,
	Run: func(cmd *cobra.Command, args []string) {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		producer, err := sarama.NewAsyncProducer([]string{"mkafka1:9092"}, config)
		if err != nil {
			panic(err)
		}

		// Trap SIGINT to trigger a graceful shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		var (
			wg                          sync.WaitGroup
			enqueued, successes, errors int
		)

		wg.Add(1)
		go func() {
			defer wg.Done()
			for range producer.Successes() {
				successes++
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for err := range producer.Errors() {
				log.Println(err)
				errors++
			}
		}()

		i := 0
	ProducerLoop:
		for {
			i++
			msg := "testing 123" + strconv.Itoa(i)
			message := &sarama.ProducerMessage{Topic: "my-topic", Value: sarama.StringEncoder(msg)}
			select {
			case producer.Input() <- message:
				enqueued++

			case <-signals:
				producer.AsyncClose() // Trigger a shutdown of the producer.
				break ProducerLoop
			}
		}

		wg.Wait()

		log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
	},
}

func init() {
	rootCmd.AddCommand(kafkaProducerCmd)
}
