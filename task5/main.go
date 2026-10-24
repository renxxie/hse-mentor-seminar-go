package main
import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	maxNum := 0
	minNum := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&x)

		if i == 0 {
			maxNum = x
			minNum = x
		} else {
			if x > maxNum {
				maxNum = x
			}
			if x < minNum {
				minNum = x
			}
		}
	}

	fmt.Printf("Максимальное значение: %d\n", maxNum)
	fmt.Printf("Минимальное значение: %d\n", minNum)
}
