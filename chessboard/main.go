package main

import "fmt"

func main() {

	var (
		board string
		size  int
	)
	fmt.Println("Введите число")
	fmt.Scanf("%d\n", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 != 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	fmt.Println(board)
}
