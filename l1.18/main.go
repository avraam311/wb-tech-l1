package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// // объявляем структуру-счетчик
// type counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// // функиця для увеличения счетчика с помощью мютексов для избежания гонок горутин
// // минус данного метода в том, что горутины блокируются каждый раз при увеличении счетчика, что снижает
// // производительность в отличие от атомарных операций
// // однако атомики небезопасны в сложных операциях, где нужно блокировать несколько переменных итд,
// // то нужно использовать мютексты, чтобы избегать ошибок
// func (c *counter) increment() {
// 	c.mu.Lock()
// 	c.value++
// 	c.mu.Unlock()
// }

// // функция для получения значения счетчика
// func (c *counter) getValue() int {
// 	return c.value
// }




// объявление структуры-счетчика
type counter struct {
	value int64
}

// функция инкремента с помощью атомика
func (ac *counter) increment() {
	atomic.AddInt64(&ac.value, 1)
}

// функция получения значения счетчика с помощью атомика
func (ac *counter) getValue() int64 {
	return atomic.LoadInt64(&ac.value)
}

func main() {
	counter := &counter{}
	var wg sync.WaitGroup

	// в 100 горутинах вызываем функцию increment
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.getValue())
}
