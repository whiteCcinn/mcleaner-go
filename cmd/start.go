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
	"context"
	"github.com/spf13/cobra"
	"mcleaner/app"
	"sync"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  `启动服务指令`,
	Run: func(cmd *cobra.Command, args []string) {
		// 2个协程并行处理各自到需求

		ctx := context.Background()
		wg := sync.WaitGroup{}

		// 从kafka拉数据到redis
		wg.Add(1)
		go func(ctx context.Context) {
			defer func() { wg.Done() }()
			pioneer := app.PioneerBackground()
			pioneer.Handle(ctx)
		}(ctx)

		// 从redis拉数据下来处理逻辑
		wg.Add(1)
		go func(ctx context.Context) {
			defer func() { wg.Done() }()
			sinker := app.SinkerBackground()
			sinker.Handle(ctx)
		}(ctx)

		// todo:
		// 主协程捕捉错误，其中一个有致命问题的话，那就结束整个程序
		// 采用 select + context channel实现
		// select {}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
