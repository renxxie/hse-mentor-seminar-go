package main
import "fmt"

func main() {
	var balance float64

	fmt.Print("Введите начальный баланс: ")
	fmt.Scan(&balance)

	for {
		fmt.Println("")
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("4. Выход")

		var choice int
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Printf("Текущий баланс: %.2f\n", balance)
			case 2:
				var amount float64
				fmt.Print("Введите сумму пополнения: ")
				fmt.Scan(&amount)

				if amount > 0 {
					balance = balance + amount
					fmt.Printf("Баланс после пополнения: %.2f\n", balance)
				} else {
					fmt.Println("Сумма должна быть больше нуля")
				}
			case 3:
				var amount float64
				fmt.Print("Введите сумму снятия: ")
				fmt.Scan(&amount)

				if amount > balance {
					fmt.Println("Недостаточно средств на балансе")
				} else if amount > 0 {
					balance = balance - amount
					fmt.Printf("Баланс после снятия: %.2f\n", balance)
				} else {
					fmt.Println("Сумма должна быть больше нуля")
				}
			case 4:
				fmt.Println("Спасибо за операцию!")
				return
			default:
				fmt.Println("Такого пункта нет, выберите число от 1 до 4")
		}
	}
}
