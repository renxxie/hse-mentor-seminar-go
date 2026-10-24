package main
import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	switch op {
		case "+":
			fmt.Printf("Результат: %g + %g = %g\n", a, b, a+b)
		case "-":
			fmt.Printf("Результат: %g - %g = %g\n", a, b, a-b)
		case "*":
			fmt.Printf("Результат: %g * %g = %g\n", a, b, a*b)
		case "/":
			if b == 0 {
				fmt.Println("Ошибка: деление на ноль невозможно")
			} else {
				fmt.Printf("Результат: %g / %g = %g\n", a, b, a/b)
			}
		default:
			fmt.Println("Ошибка: такой операции нет")
	}
}
