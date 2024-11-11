package OOP

import "fmt"

func Test_OOP() {
	a := NewAccount("Маша")
	a.SetBalance(100)

	fmt.Println(a.GetBalance()) // 100

	err := a.Deposit(10)
	if err != nil {
		panic(err) // вывод сообщения об ошибке с её описанием
	}

	err = a.Withdraw(100)
	if err != nil {
		panic(err) // вывод сообщения об ошибке с её описанием
	}

	fmt.Println(a.GetBalance()) // 10

	// создание содержимого склада
	storage := &Storage{
		Products: make(map[int]Product),
	}

	// добавление товаров на склад
	err = storage.AddProduct(Product{ID: 1, Name: "Ноутбук", Quantity: 10})
	if err != nil {
		fmt.Print("Ошибка при добавлении товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 10}

	// изменение количества товаров на складе
	err = storage.UpdateQuantity(1, -5) // продали 5 ноутбуков
	if err != nil {
		fmt.Print("Ошибка при изменении количества товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 5}
}
