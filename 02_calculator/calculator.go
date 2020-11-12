package main

import "fmt"

func main() {
	var a int /*
		# var [ var_name ] [ type ]
			1 type
				ㄱ int : 정수 // int, int8, int16, int32, int64, uint8, ...
				ㄴ float : 실수 // float, float32, float64
					# float32는 유효숫자가 7개 밖에 안됨
					# float64도 15개 밖에는 안됨
				ㄷ string : 문자열
					# 길이 : 영문자 및 숫자는 1byte, 다른 문자들은 2byte ~ 3byte
				ㄹ bool : true or false
				ㅁ rune : 문자 // 1~3byte
					# string : []rune
				...
	*/
	var b int
	a = 3 /*
		# var_name = [ value ]
		# int로 선언된 a에 실수나 문자열을 대입할 경우 에러
	*/
	b = 4
	fmt.Print(a + b)
}
