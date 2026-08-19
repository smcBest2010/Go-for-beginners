package main

import (
	"fmt"
	"struct/modules"
) // 구조체 만드는 방식, type <구조체명> struct { 객체명 타입 }

func main() {
	normal_zombie := modules.Monster{Name: "좀비", Hp: 30, Damage: []int{5, 6, 7}}
	myPlayer := modules.Player{Name: "홍길동", Hp: 40, Damage: 6}

	fmt.Printf("당신은 %s을(를) 만났다!\n", normal_zombie.Name)
	myPlayer.Hp -= normal_zombie.Damage[0]
	fmt.Printf("%s의 공격! %s은 %d만큼의 피해를 입었다. (남은 체력: %d)", normal_zombie.Name, myPlayer.Name, normal_zombie.Damage[0], myPlayer.Hp)
}
