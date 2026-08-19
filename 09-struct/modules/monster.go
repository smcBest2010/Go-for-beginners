package modules

type Monster struct { //@@@@@ 구조체 뿐만아니라 내부 변수(필드)를 외부에서 사용하려면 대문자로 해야한다.
	Name   string
	Hp     int
	Damage []int
}
