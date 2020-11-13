package main /*
	# golang에서는 int 값을 초기화하지 않을 경우 자동으로 0으로 초기화
	# string === []rune
		# golang의 string은 UTF-8
		# rune : 1 ~ 3 byte
		# string을 len만큼 index로 접근해서 출력하면 1byte씩 출력.
			# 한글 한 글자를 세 번에 나누어서 출력. 쓰레기 값.
			# 제대로 출력하고 싶으면 rune 배열로 변환 후 출력해야 함.

*/

import "fmt"

func main() {
	var A [10]int /*
		# var [ array_name ] "["[ array_len ]"]"[ type ]
	*/
	for i := 0; i < len(A); i++ { // len : 내장 키워드. 배열 길이 반환.
		A[i] = i * i
	}
	fmt.Println(A)

	s := "Hello 월드"
	s2 := []rune(s)

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]))
	}
}
