package main

import "fmt"

func main(){
	// map 初始化方式
	// 1.0 make
	m := make(map[string]int)
	// 2.0
	mm := map[string]int{}

	fmt.Println("mm:",mm)
	m["k1"]=7
	m["k2"]=13

	fmt.Println("map:",m)

	v1:=m["k1"]
	fmt.Println("v1:",v1)
	fmt.Println("len:",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	_,prs :=m["k2"]
	fmt.Println("prs:",prs)
	fmt.Println("prs:",m["k2"])
	// 如果k2存在，就返回那个值，如果不存在，返回0值，也就是说，根据这个value的类型，返回缺省值，比如string，就返回“”，int 就返回0

	n:=map[string]int{"foo":1, "bar":2}
	fmt.Println("map:",n)
}