package main

import "fmt"

func Add(a int, b int) int {
	result := a + b
	return result
}

func ArithmeticCalc(a int, b int) (int, int, float64, float64) {
	add := a + b
	sub := a - b
	multi := float64(a) * float64(b)
	divide := float64(a) / float64(b)
	return add, sub, multi, divide
}

func ArithmeticCalcWithName(a int, b int) (add int, sub int, multi float64, divide float64) {
	add = a + b
	sub = a - b
	multi = float64(a) * float64(b)
	divide = float64(a) / float64(b)
	return
}

func main() {
	fmt.Println("--- 加算 ---")
	{
		result := Add(3, 5)
		fmt.Println(result)
	}

	fmt.Println("--- 四則演算 ---")
	{
		add, sub, multi, divide := ArithmeticCalc(3, 5)
		fmt.Println("add    : ", add)
		fmt.Println("sub    : ", sub)
		fmt.Println("multi  : ", multi)
		fmt.Println("divide : ", divide)
	}

	fmt.Println("--- 四則演算（名前付き） ---")
	{
		add, sub, multi, divide := ArithmeticCalcWithName(3, 5)
		fmt.Println("add    : ", add)
		fmt.Println("sub    : ", sub)
		fmt.Println("multi  : ", multi)
		fmt.Println("divide : ", divide)
	}
}
