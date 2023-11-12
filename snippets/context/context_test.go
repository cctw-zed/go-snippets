package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// 创建一个带取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 在另一个 Goroutine 中，等待 3 秒后取消上下文
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	// 执行一个长时间运行的操作，该操作将在上下文被取消时终止
	doWork(ctx)
}

func doWork(ctx context.Context) {
	fmt.Println("Starting work...")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func TestWithTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		time.Sleep(10 * time.Second)
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
