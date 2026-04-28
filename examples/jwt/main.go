package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, err := xjwt.CreateToken(1, 1, true, 1000)
	if err != nil {
		fmt.Println("creat token failed: ", err)
		return
	}
	fmt.Println(token.Token)
}
