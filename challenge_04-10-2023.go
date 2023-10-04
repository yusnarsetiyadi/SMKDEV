package main

import (
	"fmt"
	"strconv"
)

func checkCats(catsTuti, catsNining []int, lb string) {
	var fullArr []int
	fullArr = catsTuti[1:3]
	fullArr = append(fullArr, catsNining...)
	fmt.Println(lb)
	for i := 0; i < len(fullArr); i++ {
		if fullArr[i] >= 3 {
			fmt.Printf("Kucing nomor %d adalah kucing dewasa dan berusia %d", i+1, fullArr[i])
			fmt.Println()
		} else {
			fmt.Printf("Kucing nomor %d masih anak-anak", i+1)
			fmt.Println()
		}
	}
	fmt.Println()
}

func main() {
	// basic challenge
	fmt.Println("Basic Challenge:")
	data1Tuti := []int{3, 5, 2, 12, 7}
	data1Nining := []int{4, 1, 15, 8, 3}
	data2Tuti := []int{9, 16, 6, 8, 3}
	data2Nining := []int{10, 5, 6, 1, 4}
	checkCats(data1Tuti, data1Nining, "Data Uji 1")
	checkCats(data2Tuti, data2Nining, "Data Uji 2")
	fmt.Print("\n\n")

	// advanced challenge
	var input string
	fmt.Println("Advanced Challenge:")
	fmt.Println("Input number: ")
	fmt.Scanln(&input)
	num, _ := strconv.Atoi(input)
	for i := 0; i < num; i++ {
		fmt.Printf("Current Number is : %d and the cube is %d", i+1, (i+1)*(i+1)*(i+1))
		fmt.Println()
	}
}
