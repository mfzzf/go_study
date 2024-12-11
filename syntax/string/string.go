package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("你好"))
	fmt.Println(utf8.RuneCountInString("你好"))
	fmt.Println(utf8.RuneCountInString("你好hello"))

	splited_str := strings.Split("abvc dasjk da", " ")
	for key, value := range splited_str {
		fmt.Printf("%v - > %v \n", key, value)
	}

}
