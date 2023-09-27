package test

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"server/core/gogroup"
	"testing"
	"time"
)

func TestGogroup(t *testing.T) {
	goGroup, ctx := gogroup.WithContext(context.Background())
	goGroup.SetLimit(5)

	for i := 1; i < 2; i++ {
		x := i
		goGroup.Go(func() error {
			fmt.Printf("完成任务%v: %s\n", x, time.Now().Format(time.RFC3339))
			//time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			//panic("panic is error")
			return fmt.Errorf(`%v`, "error is test")
		})
		fmt.Println("goroutine number:", runtime.NumGoroutine())
	}

	if err := goGroup.Wait(); err != nil {
		fmt.Println("err::::::", err)
	} else {
		fmt.Printf("整个大任务完成：%s\n", time.Now().Format(time.RFC3339))
	}

	fmt.Println("ctx:::::::::", ctx.Err())
	fmt.Println("ctx:::::::::", context.Cause(ctx))

	if err := gogroup.DoGo(func() error {
		fmt.Printf("完成任务(DoGoroutine): %s\n", time.Now().Format(time.RFC3339))
		var m map[int]int
		m[1] = 1
		return errors.New("asdadfadsf")
	}); err != nil {
		fmt.Println("ctx-err:::::::::", err)
	}

}
