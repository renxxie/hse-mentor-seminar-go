package main
import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	sum := 0
	maxNum := 0
	minNum := 0
	even := 0
	odd := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&x)

		sum = sum + x

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

		if x%2 == 0 {
			even = even + 1
		} else {
			odd = odd + 1
		}
	}

	fmt.Printf("Сумма всех чисел: %d\n", sum)

	if n > 0 {
		avg := float64(sum) / float64(n)
		fmt.Printf("Среднее значение: %.2f\n", avg)
	}

	fmt.Printf("Максимальное значение: %d\n", maxNum)
	fmt.Printf("Минимальное значение: %d\n", minNum)
	fmt.Printf("Количество чётных чисел: %d\n", even)
	fmt.Printf("Количество нечётных чисел: %d\n", odd)
}
