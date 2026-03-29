package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func counterStop() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan struct{})
	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			fmt.Println("горутина работает")
			time.Sleep(time.Second)
			if i == 5 {
				stopCh <- struct{}{}
				return
			}
		}
	}()
	<-stopCh
	wg.Wait()
}

func channelStop() {
	var wg sync.WaitGroup
	wg.Add(1)
	quit := make(chan string)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("горутина работает")
			}
		}
	}()
	time.Sleep(time.Second * 3)
	quit <- "quit"
	wg.Wait()
	close(quit)
}

func contextStop() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("горутина работает")
			}
		}
	}()
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}

func goexitStop() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			fmt.Println("горутина работает")
			stopCh <- struct{}{}
			runtime.Goexit()
		}
	}()
	<-stopCh
	wg.Wait()
}

func atomicStop() {
	var flag atomic.Bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for !flag.Load() {
			time.Sleep(time.Second)
			fmt.Println("горутина работает")
		}
	}()
	time.Sleep(time.Second * 3)
	flag.Store(true)
	wg.Wait()
}

func main() {
	fmt.Println("--- Метод 1: сигнал через канал после счётчика ---")
	counterStop()

	fmt.Println("\n--- Метод 2: канал quit ---")
	channelStop()

	fmt.Println("\n--- Метод 3: context cancel ---")
	contextStop()

	fmt.Println("\n--- Метод 4: runtime.Goexit ---")
	goexitStop()

	fmt.Println("\n--- Метод 5: atomic flag ---")
	atomicStop()
}
