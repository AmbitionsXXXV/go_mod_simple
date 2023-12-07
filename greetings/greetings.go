package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello函数接收一个字符串name，并返回一个问候语和一个错误值。
// 如果提供的name为空，函数返回一个错误。
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// 使用随机的问候格式。
	message := fmt.Sprintf(randomFormat(), name)
	// test failed to load greeting message
	// message := fmt.Sprintf(randomFormat())

	return message, nil
}

// Hellos函数接受一个字符串切片names，并返回一个映射（map），该映射将每个名字与其对应的问候语相关联。
// 如果在生成问候语时出现错误，函数将返回错误。
func Hellos(names []string) (map[string]string, error) {
	// 创建一个空的map来存储名字和对应的问候语。
	// map在Go中是一种键值对的集合。
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			// 如果Hello函数返回错误，则提前返回错误。
			return nil, err
		}

		// 将生成的问候语与名字相关联并存储在map中。
		messages[name] = message
	}
	return messages, nil
}

// randomFormat函数返回一个随机的问候语格式。
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 随机选择一个格式并返回。
	return formats[rand.Intn(len(formats))]
}
