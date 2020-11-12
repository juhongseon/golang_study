package main

import "fmt"

func main() {
	x := 33

	switch {
	case x == 31:
		fmt.Println("x = 31")
	case x == 32:
		fmt.Println("x = 32")
	case x == 33:
		fmt.Println("x = 33")
	} // switch 옆에 값을 안 넣을 경우 아래 case문을 돌면서 맨 처음 true의 문장을 수행
}
