// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		i := 0
// 		for {
// 			i++
// 			fmt.Println("горутина работает")
// 			// остановка горутины по условию
// 			if i == 5 {
// 				return
// 			}
// 		}
// 	}(&wg)
// 	wg.Wait()
// }





// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	quit := make(chan string)
// 	go func(*sync.WaitGroup) {
// 		defer wg.Done()
// 		// выводим текст, пока канал quit не получит сообщение и горутина не будет завершена
// 		for {
// 			select {
// 			case <-quit:
// 				return
// 			default:
// 				time.Sleep(time.Second)
// 				fmt.Println("горутина работает")
// 			}

// 		}
// 	}(&wg)
// 	time.Sleep(time.Second * 3)
// 	quit <- "quit"
// 	wg.Wait()
// 	close(quit)
// }





// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func(*sync.WaitGroup) {
// 		defer wg.Done()
// 		// выводим текст, пока не откроется канал, который дает функция ctx.Done
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				time.Sleep(time.Second)
// 				fmt.Println("горутина работает")
// 			}
// 		}
// 	}(&wg)
// 	time.Sleep(time.Second * 3)
// 	//отменяем контекст
// 	cancel()
// 	wg.Wait()
// }





// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(*sync.WaitGroup) {
// 		defer wg.Done()
// 		for {
// 			time.Sleep(time.Second)
// 			fmt.Println("горутина работает")
// 			time.Sleep(time.Second)
// 			fmt.Println("горутина работает")
// 			// завершаем горутину, текст выводится только 2 раза
// 			runtime.Goexit()
// 			time.Sleep(time.Second)
// 			fmt.Println("горутина работает")
// 		}
// 	}(&wg)
// 	wg.Wait()
// }





package main

import (
	"fmt"
	"sync"
	// пакет с атомиками для предотвращения гонок горутин
	"sync/atomic"
	"time"
)

func main() {
	// создаем флаг типа atomic.Bool, чтобы избежать гонки горутин
	// хотя в данном коде это и не возможно, так как есть только 1 горутина
	// даже если есть несколько горутин, они должны как-то влиять на флаг, чтобы нарушить работу
	// но это хороший тон в любом случае
	var flag atomic.Bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func(*sync.WaitGroup) {
		defer wg.Done()
		// пока значение !flag.Load() не равно false, горутина работает
 		for !flag.Load() {
			time.Sleep(time.Second)
			fmt.Println("горутина работает")
		}
	}(&wg)
	time.Sleep(time.Second * 3)
	// через 3 секунды меняем значение флага и горутина завершает свою работу
	flag.Store(true)
	wg.Wait()
}
