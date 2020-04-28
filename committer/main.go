package main

// TODO 拆分 peer
// 分离Committer和Endorser
//Committer提供验证区块、维护账本的功能，剔除背书、维护状态数据库的功能。
//Endorser提供模拟交易、背书功能，剔除验证区块、维护账本的功能。其状态数据库直接从Committer的区块中更新。
func main() {
	println("committer")
}
