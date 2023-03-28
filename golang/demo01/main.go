package main

import "fmt"

func main() {

	/**
	定义数组联众方式
	*/

	//arr 名称 [长度]数据类型

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	// 数组定义并初始化  下面示例等同于 var arr2 [3]int = [3]int{1,2,3}

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr2)

	// 数组定义并初始化  下面示例等同于 var arr2 [3]int = [...]int{1,2,3}  前面 [3]int的中括号里的数字取决于你后面大括号{}定义的数量

	arr3 := [...]int{1, 0}

	fmt.Printf("%T", arr3)
}
