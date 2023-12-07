package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// 设置日志的前缀，这将在所有日志消息之前显示。
	log.SetPrefix("greetings: ")
	// 设置日志的标志。这里设置为0表示无特殊标志，即仅打印日志消息，不包括日期时间等信息。
	log.SetFlags(0)

	// 创建一个包含多个名字的切片。
	names := []string{"Aimyon", "Taka", "Ryota"}

	// 调用greetings包的Hellos函数，传入名字切片。
	// 此函数将返回一个map，其中包含每个名字对应的问候语。
	messages, err := greetings.Hellos(names)
	// 如果在调用Hellos函数时出现错误，将记录并终止程序。
	if err != nil {
		log.Fatal(err)
	}

	// 打印返回的每个名字对应的问候语。
	fmt.Println(messages)
}
