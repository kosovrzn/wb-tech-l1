package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("1) Return on internal condition")
	stopByCondition()

	fmt.Println("\n2) Notification channel")
	stopByNotificationChannel()

	fmt.Println("\n3) Context cancellation")
	stopByContextCancel()

	fmt.Println("\n4) Context timeout")
	stopByContextTimeout()

	fmt.Println("\n5) runtime.Goexit")
	stopByGoexit()

	fmt.Println("\n6) Closing work channel")
	stopByClosingChannel()

	fmt.Println("\n7) Panic with recover")
	stopByPanicRecover()

	fmt.Println("\n8) OS signal via signal.Notify")
	stopBySignal()

	fmt.Println("\n9) Context deadline")
	stopByContextDeadline()

	fmt.Println("\n10) Atomic flag guard")
	stopByAtomicFlag()
}

func stopByCondition() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		const limit = 3
		for i := 1; i <= limit; i++ {
			fmt.Printf("  iteration %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}

		fmt.Println("  reached exit condition, goroutine returned")
	}()

	wg.Wait()
}

func stopByNotificationChannel() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("  got stop signal from channel")
				return
			default:
				fmt.Println("  still working...")
				time.Sleep(40 * time.Millisecond)
			}
		}
	}()

	time.Sleep(120 * time.Millisecond)
	close(stop)
	wg.Wait()
}

func stopByContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("  context canceled, stopping")
				return
			default:
				fmt.Println("  handling work")
				time.Sleep(40 * time.Millisecond)
			}
		}
	}()

	time.Sleep(120 * time.Millisecond)
	cancel()
	wg.Wait()
}

func stopByContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("  context deadline reached: %v\n", ctx.Err())
				return
			default:
				fmt.Println("  doing work until deadline")
				time.Sleep(40 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func stopByGoexit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			fmt.Println("  deferred cleanup ran before goroutine exit")
			wg.Done()
		}()

		fmt.Println("  calling runtime.Goexit()")
		runtime.Goexit()
	}()

	wg.Wait()
	fmt.Println("  goroutine stopped thanks to runtime.Goexit()")
}

func stopByClosingChannel() {
	var wg sync.WaitGroup
	tasks := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range tasks {
			fmt.Printf("  processing task %d\n", task)
			time.Sleep(40 * time.Millisecond)
		}
		fmt.Println("  task channel closed, goroutine exited loop")
	}()

	for i := 1; i <= 3; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}

func stopByPanicRecover() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  recovered panic: %v\n", r)
			}
			wg.Done()
		}()

		fmt.Println("  triggering panic to stop goroutine")
		panic("forced shutdown")
	}()

	wg.Wait()
	fmt.Println("  controller continued after panic thanks to recover")
}

func stopBySignal() {
	var wg sync.WaitGroup
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh:
				fmt.Println("  received OS signal, stopping gracefully")
				return
			default:
				fmt.Println("  running until signal arrives")
				time.Sleep(40 * time.Millisecond)
			}
		}
	}()

	time.Sleep(120 * time.Millisecond)
	// simulate an incoming SIGINT (a real app would get it from the OS)
	sigCh <- os.Interrupt
	wg.Wait()
}

func stopByContextDeadline() {
	deadline := time.Now().Add(120 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if ctx.Err() != nil {
				fmt.Printf("  deadline hit at %s: %v\n", deadline.Format(time.StampMilli), ctx.Err())
				return
			}
			fmt.Println("  waiting for fixed deadline")
			time.Sleep(40 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func stopByAtomicFlag() {
	var wg sync.WaitGroup
	var stopped atomic.Bool
	stop := make(chan struct{})

	wg.Add(2)

	// worker
	go func() {
		defer wg.Done()
		for {
			if stopped.Load() {
				fmt.Println("  worker observed atomic stop flag")
				return
			}
			select {
			case <-stop:
				fmt.Println("  worker draining stop channel")
				return
			default:
				fmt.Println("  worker ticking")
				time.Sleep(30 * time.Millisecond)
			}
		}
	}()

	// two competing stoppers; only one should win and trigger cleanup
	go func() {
		defer wg.Done()
		time.Sleep(90 * time.Millisecond)
		if stopped.CompareAndSwap(false, true) {
			close(stop)
			fmt.Println("  stopper A set flag and closed channel")
		} else {
			fmt.Println("  stopper A saw flag already set")
		}
	}()

	time.Sleep(120 * time.Millisecond)
	if stopped.CompareAndSwap(false, true) {
		close(stop)
		fmt.Println("  stopper B set flag and closed channel")
	} else {
		fmt.Println("  stopper B noticed flag already set")
	}

	wg.Wait()
}
