package main

import (
	"fmt"

	"github.com/suwakei/go-zhconv/zhconv"
)

func main() {
    fmt.Println(zhconv.Z2h("Ｈello World!")) // "Hello, world!"
    fmt.Println(zhconv.H2z("Hｅｌｌｏ， ｗｏｒｌｄ！")) // "Ｈｅｌｌｏ， ｗｏｒｌｄ！"
}