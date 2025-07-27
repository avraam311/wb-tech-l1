// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	get := make(chan int)
// 	var n int
// 	fmt.Scan(&n)
// 	// присвоение переменной timeOut канала, возвращаемого функцией time.After
// 	// в этом канале одно значение для чтения - текущее время
// 	timeOut := time.After(time.Second * time.Duration(n))

// 	// постоянная запись в канал get с помощью бесконечного цикла
// 	// цикл запущен в анонимной горутине, чтобы основная горутина могла перейти к запуску функции reader
// 	go func() {
// 		for {
// 			get <- 1
// 		}
// 	}()

// 	for {
// 		select {
// 		case <-get:
// 			fmt.Println("прочитано из канала get")
// 			time.Sleep(time.Millisecond * 500)
// 		// чтение из канала timeOut, чтобы по истечении времени завершить горутину
// 		case <-timeOut:
// 			fmt.Println("время истекло")
// 			return
// 		}
// 	}
// }





// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	get := make(chan int)
// 	var n int
// 	fmt.Scan(&n)
// 	// присвоение переменной timeOut таймера(структура, в которую встроен канал C, из которого можно прочитать
// 	// текущее время)
// 	timeOut := time.NewTimer(time.Second * time.Duration(n))

// 	// постоянная запись в канал get с помощью бесконечного цикла
// 	// цикл запущен в анонимной горутине, чтобы основная горутина могла перейти к запуску функции reader
// 	go func() {
// 		for {
// 			get <- 1
// 		}
// 	}()

// 	for {
// 		select {
// 		case <-get:
// 			fmt.Println("прочитано из канала get")
// 			time.Sleep(time.Millisecond * 500)
// 		// чтение из канала timeOut.C, чтобы по истечении времени завершить горутину
// 		case <-timeOut.C:
// 			fmt.Println("время истекло")
// 			return
// 		}
// 	}
// }





package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	get := make(chan int)
	var n int
	fmt.Scan(&n)
	// инициализируем контекст с таймаутом, который сам отменяется по истечении времени
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(n))
	// отложенно вызываем функции cancel на случай непредвиденных ситуаций, чтобы избежать утечку памяти,
	// несмотря на то, что контекст сам отменится по истечении времени
	defer cancel()
	
	// постоянная запись в канал get с помощью бесконечного цикла
	// цикл запущен в анонимной горутине, чтобы основная горутина могла перейти к запуску функции reader
	go func() {
		for {
			get <- 1
		}
	}()

	for {
		select {
		case <-get:
			fmt.Println("прочитано из канала get")
			time.Sleep(time.Millisecond * 500)
		// чтение из канала возвращаемого функцией ctx.Done
		// изначально этот канал закрыт, а значит мы никогда не достигнем return
		// однако, когда истекает время действия контекста, этот канал открывается и отправляет пустой
		// объект типа struct{}, следовательно, мы читаем из канала и выполняем блок case с return
		case <-ctx.Done():
			fmt.Println("время истекло")
			return
		}
	}
}
