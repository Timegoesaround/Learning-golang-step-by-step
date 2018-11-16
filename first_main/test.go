package main

import "fmt"

func main() {
	fmt.Println("hello")
	// quiz1()

	quiz2()
}

/*
 * 创建一个基于for 的简单的循环。使其循环10 次，并且使用fmt 包打印出计数
 * 器的值。
 */
func quiz1() {
	a := 0
	//using for
	// for i := 0; i < 10; i++ {
	// 	a += i
	//fmt.Println(a)
	// }

	//using flag and goto
	i, a := 0, 0
begin:
	a += i
	i++
	fmt.Println(a)
	if i < 10 {
		goto begin
	}

	var arr [10]int
	a2 := 0
	for index := 0; index < 10; index++ {
		a2 += index
		arr[index] = a2
		fmt.Println(a2)
	}

}
