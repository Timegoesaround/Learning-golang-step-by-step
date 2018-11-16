/*go 学习笔记
 *
 * 包名小写，不应当有下划线
 * 包名为main表明 该包为可执行函数 build后生成 可执行文件
 *
 */
package main

/*
 * 导入包
 * import (
	 "fmt"
	 "**"
 )
 *可以给导入的包取别名
  import bar "bytes" //bar 为别名
*/
import (
	"fmt"
)

/*
 * 函数命名
 *  共有函数 大写字母开头 可导出，可在包的外部调用。
 *  私有函数 小写字母开头，包外不可用
 *  测试函数 TestXxx(t *testing.T) 使用 go test可执行测试函数
 *
 */
func main() {
	/*
	 * 变量声明后必须要使用，不然编译报错
	 */
	var b bool
	b = false

	a := 12

	/*
	 *	常量即使未被使用，也不会报错
	 */
	const (
		yes = 0
		no  = 1
	)

	/*
	 * 字符串操作
	 */
	s := "string head" +
		"string tail"
	/*
	 * 原始字符串符号内不转义
	 */
	s1 := `string head
	 string tai/sl`
	fmt.Print("b 先声明，再赋值")
	fmt.Println(b)
	fmt.Printf("a := %d", a)
	fmt.Printf("hello, world\n")

	fmt.Printf("小心字符拼接符 '+' 放在每一行末尾" + s)
	fmt.Printf("可以使用原始字符串符号``" + s1)

	fmt.Println()
	//此处 左大括号必须与 if处于同一行
	if flag := 0; flag == yes {
		fmt.Println("true" + (string)(yes))
	} else {
		fmt.Println("false")
	}

	//go中 有 goto ,还需要for循环嘛？(估计是嫌弃goto写循环太麻烦了),当然也更容易死循环 0.0
	i := 1
FlagDead:
	println(i)
	i++
	if i <= 100 {
		goto FlagDead
	}
Outter:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if j > 50 {
				break Outter //跳出外层循环
			}
			// fmt.Println(j)
		}
	}
	//好用的迭代器
	list := []string{"h", "e", "l", "l", "o", ",", "world"}

	for k, v := range list {
		fmt.Printf("%d:%s\n", k, v)
	}

	/*
	 * switch 可以是常量，整数，默认为bool
	 * switch 一个分支有多个条件时，用 ','分割
	 */
	switch {

	case false:
		fmt.Println("false")
		fallthrough //fallthrough 可以使程序匹配失败后继续向下尝试
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default")
		array()
	}

	practice()
}

/*
 * go 中有少数内部函数，可使用 go doc builtin查看
 * close new panic complex
 * delete make recover real
 * len append print imag
 * cap copy println
 */

/*
 * 数组类型  [n]<type>
 * 数组之间赋值 传递的是一个副本，而非指针
 *
 */
func array() {
	var array [3]int
	array[0] = 0
	array[1] = 1
	array[2] = 2

	//或
	// array := [3]int{1, 2, 3}
	//或
	// array := [...]int{1,2,3}

	for k, v := range array {
		fmt.Printf("the number in array[%d] is %d", k, v)
	}

	/*
	 * 声明多维数组时，有些内容必须给定
	 */
	array2 := [3][2]int{[...]int{1, 2}, [...]int{2, 3}, [...]int{3, 4}}
	//array2 := [3][2]int{{1, 2}, {2, 3}, {3, 4}} 也行

	for k, v := range array2 {
		fmt.Printf("index%d", k)
		for vk, v2 := range v {
			fmt.Printf("index %d is %d", vk, v2)
		}
	}

	/*
	 * slice 为可变长array， 是一个指向array的指针
	 * slice 总是和一个固定长度的array成对出现
	 * 注： 引用类型使用make 创建
	 */
	sl := make([]int, 10)

	fmt.Printf("新建slice长度为：%d", len(sl))

	/*
	 * copy(dst,src) 从源slice src复制元素到目标 dst，并且返回复制的元素个数
	 * 复制的数量是 src 和 dst中较小的一个
	 *
	 */

}

/*TestHello 测试函数

 */
func TestHello() {
	return
}
