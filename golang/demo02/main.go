package main

import "fmt"

func main() {

	/**

	创建map的多种方法  map创建需要主要需要定义 不然是会报错的


	*/

	/**
	  这种方式是最鸡肋的
	*/
	var mp map[string]int

	mp = map[string]int{"hello": 1, "go": 1}

	fmt.Println(mp)

	var mp2 = map[string]int{"hello": 1, "go": 2}

	fmt.Println(mp2)

	/**
	这种是最常用的定义方式 也比较通俗易懂
	*/
	mp3 := make(map[string]int)

	mp3["hello"] = 1

	fmt.Println(mp3)

}
