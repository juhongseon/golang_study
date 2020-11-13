package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func getRand0to9() int {
	num, _ := rand.Int(rand.Reader, big.NewInt(9))
	return int(num.Int64()) + 1
}

func getOrigin() [3]int {
	var res [3]int
	res[0] = getRand0to9()
	for {
		temp := getRand0to9()
		if res[0] != temp {
			res[1] = temp
			break
		}
	}
	for {
		temp := getRand0to9()
		if res[0] != temp && res[1] != temp {
			res[2] = temp
			break
		}
	}
	return res
}

func getUserInput() [3]int {
	var res [3]int
	fmt.Print("input : ")
	var raw int
	for {
		_, err := fmt.Scanf("%d\n", &raw)
		if err != nil {
			fmt.Println("wrong input")
			continue
		}
		res[0] = raw / 100
		res[1] = (raw % 100) / 10
		res[2] = raw % 10
		break
	}
	return res
}

func compare(origin [3]int, user [3]int) string {
	var strike int
	var ball int

	for u := 0; u < 3; u++ {
		for o := 0; o < 3; o++ {
			if user[u] == origin[o] {
				if u == o {
					strike++
				} else {
					ball++
				}
			}
		}
	}
	return strconv.Itoa(ball) + "B" + strconv.Itoa(strike) + "S"
}

func main() {
	origin := getOrigin()
	user := getUserInput()
	try := 1
	for {
		res := compare(origin, user)
		if res == "0B3S" {
			fmt.Println("success. try : ", try)
			break
		} else {
			fmt.Println(res)
			user = getUserInput()
			try++
		}
	}
}
