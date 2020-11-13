package main

import "fmt"

func main() {
	var a int
	var p *int // 포인터 선언. 모든 타입(구조처 포함)에 대해 가능.
	p = &a     // 주소값 대입.

	a = 3
	fmt.Println(a, p, *p) // 포인터를 통해 해당 주소의 값에 접근.
}
