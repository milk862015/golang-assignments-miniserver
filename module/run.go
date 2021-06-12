package module

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// Run  运行
func Run() {
	// 上下文
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	// 初始化
	srv := &http.Server{Addr: ":8888"}
	signalChannel := NewSignalChannel()

	// 启动
	group.Go(func() error {
		return RunHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		// 关闭服务器
		fmt.Println("[mini]http server stop")
		err := srv.Shutdown(errCtx)
		if err != nil {
			return errors.New("[mini-error] server shutdown err:" + err.Error())
		}
		return nil
	})

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				fmt.Println("[mini]run in for by errCtx.Done")
				return errCtx.Err()
			case <-signalChannel:
				fmt.Println("[mini]signal channel")
				cancel()
				return errors.New("[mini-error]signal cancel")
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("[mini] group error:")
		fmt.Println(err)
		return
	}

	fmt.Println("[mini]done")
}
