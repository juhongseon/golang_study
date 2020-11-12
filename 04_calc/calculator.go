package main

import (
	"bufio" // 한 줄 입력받기
	"fmt"
	"os"      // 표준 입력
	"strconv" // 문자열 숫자로 바꿔줌
	"strings" // 입력 값 제어
) // 표준 패키지는 공식 문서로 공부 필요

func main() {
	fmt.Print("input 1 : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n') // 개행 문자가 나올 때까지 읽겠다는 의미
	line = strings.TrimSpace(line)
	n1, _ := strconv.Atoi(line) // ASCII to integer

	fmt.Print("input 2 : ")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Printf("inputs = %d, %d\n", n1, n2)

	fmt.Print("operator : ")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	} else {
		fmt.Printf("wrong operator\n")
	}

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	default:
		fmt.Printf("wrong operator\n")
	} // break 없어도 됨
}
