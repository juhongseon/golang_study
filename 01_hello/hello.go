package main /*
	# package [ package_name ] : 패키지 선언 필수
		1 package (at golang) : 라이브러리 선언
		2 package_name
			# main : 프로그램의 시작점
			# main으로 선언 시 파일 내부에 func main이 있어야 함
				function main is undeclared in the main package
			# main이 아닌 패키지 run 할 경우 에러메시지
				go run: cannot run non-main package
	# 같은 디렉토리 내에 main package가 복수 존재할 경우 에러
*/

import "fmt" /*
	# import "[ package_name ]" : 해당 패키지의 함수들을 사용하겠다는 선언
	# 여러 패키지를 사용하고 싶을 경우
		import {
			"fmt"
			"sort"
			...
		}
*/

func main() { /*
		# func [ func_name ]([ input ]) [output] {
			# go lang은 중괄호 위치가 항상 위와 같아야 함
		}
			# input, output은 각각 없을 경우 생략
	*/
	fmt.Print("Hello, Go.")
}
