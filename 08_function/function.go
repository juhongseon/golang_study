package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) { /*
		# 타입 같으면 생략 가능
		# 복수 리턴값 가능
	*/
	return y, x
}

func fbnc(end int) int {
	if end < 2 {
		return 1
	}
	return fbnc(end-1) + fbnc(end-2)
}

func main() {
	fmt.Println(add(1, 2)) // golang은 항상 값이 복사됨. 참조 X

	a, b := swap(2, 3)
	fmt.Println(a, b)

	fmt.Println(fbnc(10))
}
