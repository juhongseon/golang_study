package main /*
	golang은 모든 반복문을 for만으로 작성. 다른 것은 없다.
	break, continue는 다른 언어와 동일.
*/

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("last i = ", i)
	fmt.Println("========")

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	// fmt.Println("last j = ", j)
	// 위 문장은 에러. 스코프에 주의
	fmt.Println("========")

	// 아래와 같이 작성하면 가능
	var j int
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	fmt.Println("last j = ", j)
	fmt.Println("========")

	// for문의 조건문을 생략하는 경우 조건에 true가 들어감. 무한 루프.
}
