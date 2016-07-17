package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//SIGKILL SIGSTOP can't be ignored
func main() {
	fmt.Println("---------------signal--------------")
	sigRecv := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}

	signal.Notify(sigRecv, sigs...)

	for s := range sigRecv {
		fmt.Printf("Receive signal %s\n", s)
		signal.Stop(sigRecv)
		close(sigRecv)
	}

	time.Sleep(time.Second * 10)
	fmt.Println("process exit")
}
