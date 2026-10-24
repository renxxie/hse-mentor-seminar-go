package main
import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	positive := 0
	negative := 0
	zeros := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&x)

		if x > 0 {
			positive = positive + 1
		} else if x < 0 {
			negative = negative + 1
		} else {
			zeros = zeros + 1
		}
	}

	fmt.Printf("Положительных чисел: %d\n", positive)
	fmt.Printf("Отрицательных чисел: %d\n", negative)
	fmt.Printf("Нулей: %d\n", zeros)
}
