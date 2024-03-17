package main

import "fmt"

func main() {
	count, total := 0, 0
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			count++
			total = total + i
		}
	}
	average := float64(total) / float64(count)
	fmt.Println(average)
	fmt.Println(float64(total) / float64(count))
}
