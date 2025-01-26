package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	svr *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		svr: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// NOTE(ブレースフルシャットダウン): シグナルを受け取ったらコンテキストをキャンセルする
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.svr.Serve(s.l); err != nil &&
			// NOTE: ErrServerClosedはサーバーが正常に終了した場合に返されるエラーのため、異常ではない
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// NOTE(ブレースフルシャットダウン): コンテキストが完了（=キャンセルシグナルを受け取るまで)ブロック
	<-ctx.Done()
	if err := s.svr.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// グレースフルシャットダウンの終了を待つ
	return eg.Wait()
}
