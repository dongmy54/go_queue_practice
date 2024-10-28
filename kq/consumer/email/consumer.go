package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var c kq.KqConf
	conf.MustLoad("config.yaml", &c)

	// 本质上kafka发送的就是一条消息 至于谁能去消费依靠这里的c决定
	q := kq.MustNewQueue(c, kq.WithHandle(func(ctx context.Context, k, v string) error {
		fmt.Printf("email消费者=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
