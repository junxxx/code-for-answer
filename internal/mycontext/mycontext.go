package mycontext

import (
	"context"
	"fmt"
	"time"
)

const (
	key = "age"
	val = "23"
)

func enrichValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func routine(ctx context.Context) {
	v := ctx.Value(key)
	fmt.Println("ctx'key is: ", v)
}

func test() {
	ctx := context.Background()
	ctx = enrichValue(ctx, key, val)
	routine(ctx)
}

func slowTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println(" runing task")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func testWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go slowTask(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
