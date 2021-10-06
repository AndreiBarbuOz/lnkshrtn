package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	var (
		args []string
	)
	fmt.Printf("Started main test\n")

	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "DEVEL"):
		case strings.HasPrefix(arg, "-test"):
		default:
			args = append(args, arg)
		}
	}

	waitCh := make(chan int, 1)

	os.Args = args
	os.Args[0] = "lnkshrtn-server"
	go func() {
		main()
		close(waitCh)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-signalCh:
		fmt.Printf("received signal")
		time.Sleep(time.Second * 5)
		return
	case <-waitCh:
		fmt.Printf("exited from server")
		return
	}
}
