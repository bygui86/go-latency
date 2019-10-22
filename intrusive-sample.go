package main

import (
	"fmt"
	"time"
)

func main() {
	expensivePrint()
	fmt.Println("Finished executing main")
}

func expensivePrint() {
	start := time.Now()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Current number is %d \n", i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Time taken by expensivePrint function is %v \n", time.Since(start))
}
