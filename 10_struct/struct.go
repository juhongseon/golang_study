package main /*
	# 구조체는 프로그램의 응집성을 높이기 위함.
*/

import "fmt"

type person struct { /*
		# type [ struct_name ] struct
			1 type : type alias
			2 struct_name을 대문자로 쓸 경우 다른 패키지에서도 참조 가능. 접근 제한자.
			3 struct // golang의 struct는 일급 객체.
	*/
	name string
	age  int
}

func (p *person) setName(name string) { /*
		# 포인터 리시버로 받지 않으면 setter의 역할을 할 수 없음.
	*/
	p.name = name
	fmt.Println("setName call...")
}

func (p person) getName() string { /*
		# 단순 조회는 value receiver로도 충분히 가능.
		# 그러나 이렇게 하면 괜히 해당 구조체만큼 메모리를 더 생성하니 포인터 리시버로 받는 게 좋을 듯 하다.
	*/
	return p.name
}

func main() {
	var p person // name == "", age == 0
	p1 := person{"jhs", 29}
	p2 := person{name: "jhs", age: 29}
	p3 := person{name: "jhs"} // age == 0

	fmt.Println(p, p1, p2, p3)

	p.setName("jhs")
	fmt.Println(p, p1, p2, p3)

	fmt.Println(p.getName())
}
