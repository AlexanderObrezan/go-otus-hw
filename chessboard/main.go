package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		board strings.Builder
		size  int
	)
	fmt.Println("Введите число")
	fmt.Scanf("%d\n", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 != 0 {
				board.WriteString("#")
			} else {
				board.WriteString(" ")
			}
		}
		board.WriteString("\n")
	}
	fmt.Println(board.String())
}
