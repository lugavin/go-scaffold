package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Context 的主要作用是提供一种标准的机制来处理请求链路中的上下文数据传递、取消和超时等问题，从而使代码更加简洁、可读性更好、可维护性更高。
func TestContext(t *testing.T) {
	ctx := context.Background()

	// 创建一个带有取消信号的 context，超时时间为 1 秒
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// 使用 WithValue 方法在 context 中传递请求范围的数据
	ctx = context.WithValue(ctx, "requestID", "12345")

	// 调用一个模拟的处理函数，传入 context
	result := processRequest(ctx)

	// 输出结果
	fmt.Println(result)
}

func processRequest(ctx context.Context) string {
	// 从 context 中获取请求 ID
	requestID := ctx.Value("requestID").(string)

	// 模拟一个耗时 2 秒的操作
	time.Sleep(2 * time.Second)

	// 判断 context 是否已经被取消
	select {
	case <-ctx.Done():
		// 如果 context 已经被取消，返回错误信息
		return fmt.Sprintf("request %s cancelled", requestID)
	default:
		// 如果 context 没有被取消，返回处理结果
		return fmt.Sprintf("request %s processed", requestID)
	}
}
