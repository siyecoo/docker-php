package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (that Person) ChangeAge1(age int) {
	that.Age = age
}

func (that *Person) ChangeAge2(age int) {
	that.Age = age
}

func main() {

	/**

	  结构体的使用

	*/

	//  有指定名称的定义
	p := Person{Name: "Siyecoo", Age: 18}

	fmt.Println("p的值:", p)

	// 不指定名称

	p1 := Person{"hello", 20}

	fmt.Println("p1的值:", p1)

	//部分指定

	p2 := Person{Age: 15}
	fmt.Println("p2的值:", p2)

	//调用方法未指定结构体指针的情况下无法更改结构的属性

	p3 := Person{Age: 16}

	p3.ChangeAge1(20)

	fmt.Println("p3未改变Age的值:", p3)

	fmt.Println(p3)

	p3.ChangeAge2(15)
	fmt.Println("p3改变Age的值:", p3)

}
