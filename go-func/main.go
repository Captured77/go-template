package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/nxadm/tail"
	"golang.org/x/sync/errgroup"
)

type TermFileError struct{}
type SocketFileError struct{}

func (m *TermFileError) Error() string {
	return "termFile got detected"
}

func (m *SocketFileError) Error() string {
	return "socketFile got removed"
}

type Log struct {
	LogFile string
	ctx     context.Context
	g       *errgroup.Group
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	g, gctx := errgroup.WithContext(ctx)
	l := Log{
		LogFile: "/var/log/kern.log",
		ctx:     gctx,
		g:       g,
	}

	g.Go(l.tailLog)

	if err := g.Wait(); err != nil {
		if !(errors.Is(err, context.Canceled) || errors.Is(err, &TermFileError{}) || errors.Is(err, &SocketFileError{})) {
			fmt.Println("exit")
			os.Exit(1)
		}
		// Exit cleanly on clean termination errors
	}
}

func (v *Log) tailLog() error {
	t, err := tail.TailFile(v.LogFile, tail.Config{
		Follow: true,
	})
	if err != nil {
		return err
	}

	for {
		select {
		case line, ok := <-t.Lines:
			if !ok {
				// log.Log.V(4).Infof("tail error: %v", line)
				fmt.Printf("err: %v", err)
			} else if line != nil {
				if line.Err != nil {
					fmt.Printf("line error: %v", err)
				} else {
					fmt.Println(line.Text)
				}
			}
		case <-v.ctx.Done():
			return v.ctx.Err()
		}
	}

}
