package main

import "fmt"

//两种方法来实现在类型中嵌入功能：
//聚合：包含一个所需功能类型的具名字段。
//内嵌：内嵌所需功能类型。

//例子：通过Log类型来包含日志功能，Log类型只是简单包含一个积累
//的消息。

//以下是聚合方式的实现

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
	// 内嵌方式
	// Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can !"

	c = &Customer{"Barak Obama", &Log{"1 - Yes we can up up !"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
