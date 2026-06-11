package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("claims.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
