package main

import "fmt"

func main() {
	a := 4 /*
		# [ var_name ] := [ value ] : 선언대입문
	*/
	b := 2

	fmt.Printf("4 & 2 = %v\n", a&b) /*
		# %v : 파라미터에 맞는 타입으로 출력하라는 뜻. 자동으로 타입 추정.
	*/
	fmt.Printf("4 | 2 = %v\n", a|b)
	fmt.Printf("4 ^ 2 = %v\n", a^b)

	a = 21
	c := a % 10
	d := a / 10
	fmt.Printf("%v = %v * 10 + %v\n", a, d, c)

	if 3 > 4 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	} /*
		# 조건식에 괄호 안씀
		# else 키워드와 중괄호 위치 다 지켜야 함
	*/
}
