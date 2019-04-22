package main

import "fmt"

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main(){
	i:=1
	fmt.Println("initial:", i)
    zeroval(i)
	fmt.Println("zeroval:", i)
	//实参和行参都是类型 i，可以接受。此时在zeroval()中的 i只是 i的值拷贝，所以zeroval()的修改影响不到t1。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}