package concurrency

import "fmt"

func squares(num int, ch chan int) {
	for i := 0; i < num; i++ {
		ch <- i * i // запись в канал
	}

	close(ch) // закрытие канала
}

func Test_ch() {
	ch := make(chan int) // инициализация канала

	go squares(5, ch)

	// чтение данных до закрытия канала
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Завершение...")
}
