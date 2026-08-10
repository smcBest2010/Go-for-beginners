package main

import (
	"badass/average"
	"fmt"
)

func main() {
	fmt.Printf("%.2f\n", average.Avg(100, 200, 200))
	fmt.Println(average.AdditionMin(10, 20, 30, 40))
}
