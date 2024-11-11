package concurrency

import (
	"fmt"
	"time"
)

func iterate(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, i)
	}
}

func Test_conc() {
	iterate("Обычный вызов:") // синхронно
	go iterate("Вызов с go:") // параллельно

	go func(str string) {
		fmt.Println(str)
	}("Анонимная функция с go")

	time.Sleep(2 * time.Second)
}
