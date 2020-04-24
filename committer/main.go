package main

import "fmt"

type abc struct {
	name string
	age  string
}

func (abc *abc) String() string {
	return abc.name + "=" + string(abc.age)
}

// TODO 拆分 peer
// 分离Committer和Endorser
//Committer提供验证区块、维护账本的功能，剔除背书、维护状态数据库的功能。
//Endorser提供模拟交易、背书功能，剔除验证区块、维护账本的功能。其状态数据库直接从Committer的区块中更新。
func main() {

	s := new(abc)
	s.name = "name"
	s.age = "10"

	fmt.Printf("%s", s)

	//fmt.Println("记账节点")
	//c := make(chan string, 2)
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	c <- "a"
	//}()
	//
	////println( <- c)
	//
	//go func() {
	//	c <- "b"
	//}()
	//
	//for i := 0; i < cap(c); i++ {
	//	println( <- c)
	//}
}
